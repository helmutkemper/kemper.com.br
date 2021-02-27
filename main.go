package main

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type UniqueKeyInterface interface {
	Make() (uniqueKey string)
}

type PasswordInterface interface {
	MakeHash(password []byte) (hash []byte, err error)
	CheckHash(password, hash []byte) (match bool)
	NewPasswordRule(password []byte) (err error)
}

type DataInterface interface {
	Init() (err error)
	Connect(filePath string) (err error)
	MailExists(mail string) (found bool, err error)
	MailHasVerified(mail string) (mailHasVerified bool, err error)
	UserInsert(name, nickname string, gender, userType int, mail, password string, mailVerified int) (err error)
	GetPassword(mail string) (password string, err error)
}

type UniqueKey struct{}

func (e *UniqueKey) Make() (uniqueKey string) {
	var uniqueKeyBase = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q",
		"r", "s", "t", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "!", "@", "#",
		"$", "%", "&", "*", "(", ")", "_", "=", "-", "+", "[", "]", "{", "}", "|", "/", "?", "<", ">", ";", ":"}
	var randonGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i != 50; i += 1 {
		uniqueKey += uniqueKeyBase[randonGenerator.Intn(82-1)]
	}

	return
}

type Password struct{}

func (e *Password) ruleLength(password []byte) (err error) {
	if len(password) < 8 {
		err = errors.New("the password must be 8 letters or more")
	}

	return
}

func (e *Password) ruleOneSpecialChars(password []byte) (err error) {
	var char []byte
	var specialChars = [][]byte{[]byte("`"), []byte("~"), []byte("!"), []byte("@"), []byte("#"), []byte("$"),
		[]byte("%"), []byte("^"), []byte("&"), []byte("*"), []byte("("), []byte(")"), []byte("-"), []byte("_"),
		[]byte("+"), []byte("="), []byte("["), []byte("{"), []byte("]"), []byte("}"), []byte("|"), []byte("\\"),
		[]byte(";"), []byte(":"), []byte("\""), []byte("'"), []byte("<"), []byte(">"), []byte(","), []byte("."),
		[]byte("/"), []byte("?")}
	for char = range specialChars {
		if bytes.Contains(char, password) {
			return
		}
	}

	err = errors.New("the password must be one special char")
	return
}

func (e *Password) ruleUpperLetter(password []byte) (err error) {
	var char []byte
	var specialChars = [][]byte{[]byte("A"), []byte("B"), []byte("C"), []byte("D"), []byte("E"), []byte("F"),
		[]byte("G"), []byte("H"), []byte("I"), []byte("J"), []byte("K"), []byte("L"), []byte("M"), []byte("N"),
		[]byte("O"), []byte("P"), []byte("Q"), []byte("R"), []byte("S"), []byte("T"), []byte("V"), []byte("W"),
		[]byte("X"), []byte("Y"), []byte("Z")}
	for char = range specialChars {
		if bytes.Contains(char, password) {
			return
		}
	}

	err = errors.New("the password must be one upper case char")
	return
}

func (e *Password) ruleLowerCase(password []byte) (err error) {
	var char []byte
	var specialChars = [][]byte{[]byte("a"), []byte("b"), []byte("c"), []byte("d"), []byte("e"), []byte("f"),
		[]byte("g"), []byte("h"), []byte("i"), []byte("j"), []byte("k"), []byte("l"), []byte("m"), []byte("n"),
		[]byte("o"), []byte("p"), []byte("q"), []byte("r"), []byte("s"), []byte("t"), []byte("v"), []byte("w"),
		[]byte("x"), []byte("y"), []byte("z")}
	for char = range specialChars {
		if bytes.Contains(char, password) {
			return
		}
	}

	err = errors.New("the password must be one lower case char")
	return
}

func (e *Password) NewPasswordRule(password []byte) (err error) {
	err = e.ruleLength(password)
	if err != nil {
		return
	}

	err = e.ruleOneSpecialChars(password)
	if err != nil {
		return
	}

	err = e.ruleUpperLetter(password)
	if err != nil {
		return
	}

	err = e.ruleLowerCase(password)
	return
}

func (e *Password) MakeHash(password []byte) (hash []byte, err error) {
	hash, err = bcrypt.GenerateFromPassword(password, 30)
	return
}

func (e *Password) CheckHash(password, hash []byte) (match bool) {
	var err = bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

type SQLite3 struct {
	database *sql.DB
}

func (e *SQLite3) Connect(filePath string) (err error) {
	e.database, err = sql.Open("sqlite3", filePath)
	return
}

func (e *SQLite3) Init() (err error) {
	err = e.userCreate()
	if err != nil {
		return
	}

	return
}

func (e *SQLite3) userCreate() (err error) {
	var statement *sql.Stmt
	statement, err = e.database.Prepare(`
		CREATE TABLE IF NOT EXISTS 
			user (
				id INTEGER PRIMARY KEY AUTOINCREMENT, 
				name TEXT,
				nickname TEXT, 
				gender INTEGER,
				userType INTEGER,
				mail TEXT,
				password TEXT,
				mailVerified INTEGER
			)
		`,
	)
	if err != nil {
		return
	}

	_, err = statement.Exec()
	return
}

func (e *SQLite3) MailExists(mail string) (found bool, err error) {
	var rows *sql.Rows
	rows, err = e.database.Query(
		`
			SELECT 
				id 
			FROM 
				user 
			WHERE 
				mail = ?
		`,
		mail,
	)
	if err != nil {
		return
	}

	found = rows.Next()
	return
}

func (e *SQLite3) MailHasVerified(mail string) (mailHasVerified bool, err error) {
	var rows *sql.Rows
	rows, err = e.database.Query(
		`
			SELECT 
				   id 
			FROM 
				 user 
			WHERE 
				  mail = ? 
			  AND mailVerified = 1
		  `,
		mail,
	)
	if err != nil {
		return
	}

	mailHasVerified = rows.Next()
	return
}

func (e *SQLite3) UserInsert(name, nickname string, gender, userType int, mail, password string, mailVerified int) (err error) {
	var statement *sql.Stmt
	statement, err = e.database.Prepare(`
		INSERT INTO 
		    user (name, nickname, gender, userType, mail, password, mailVerified) 
		VALUES   
		       	 (?,    ?,        ?,      ?,        ?,    ?,        ?)`,
	)
	if err != nil {
		return
	}

	_, err = statement.Exec(name, nickname, gender, userType, mail, password, mailVerified)
	return
}

func (e *SQLite3) GetPassword(mail string) (password string, err error) {
	var rows *sql.Rows
	rows, err = e.database.Query(`
		SELECT 
		       password 
		FROM 
		     user 
		WHERE 
		      mail = ?`,
		mail,
	)
	if err != nil {
		return
	}

	for rows.Next() {
		err = rows.Scan(&password)
		return
	}

	return
}

type BusinessRules struct {
	UniqueKey UniqueKeyInterface
	Password  PasswordInterface
	Data      DataInterface
}

//RFC 5322 Official Standard
func (e *BusinessRules) MailSyntax(mail string) (matched bool, err error) {
	matched, err = regexp.MatchString("(?:[a-z0-9!#$%&'*+=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])", mail)
	return
}

func (e *BusinessRules) UserLogin(mail, password string) (successful bool, err error) {
	var hashPassword string

	successful, err = e.Data.MailExists(mail)
	if err != nil {
		return
	}
	if successful == false {
		return false, errors.New("email not found")
	}

	successful, err = e.Data.MailHasVerified(mail)
	if err != nil {
		return
	}
	if successful == false {
		return false, errors.New("email has not yet been verified")
	}

	hashPassword, err = e.Data.GetPassword(mail)
	if err != nil {
		return
	}

	successful = e.Password.CheckHash([]byte(password), []byte(hashPassword))
	if successful != true {
		return false, errors.New("password does not match")
	}
	return
}

func (e *BusinessRules) UserNew(name, nickname string, gender int, mail, password string) (err error) {
	var pass bool

	pass, err = e.MailSyntax(mail)
	if err != nil {
		return
	}
	if pass == false {
		err = errors.New("e-mail must be a valid syntax")
		return
	}

	err = e.Password.NewPasswordRule([]byte(password))
	if err != nil {
		return
	}

	err = e.Data.UserInsert(name, nickname, gender, 0, mail, password, 0)
	return
}

func main() {
	var err error
	var database *sql.DB
	var statement *sql.Stmt
	var rows *sql.Rows

	database, err = sql.Open("sqlite3", "./db/database.sqlite")
	if err != nil {
		log.Fatalf("sql.Open().error: %v", err)
	}

	statement, err = database.Prepare(
		"CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY AUTOINCREMENT, firstname TEXT, lastname TEXT)",
	)
	if err != nil {
		log.Fatalf("sql.Prepare().error: %v", err)
	}

	_, err = statement.Exec()
	if err != nil {
		log.Fatalf("sql.Exec().error: %v", err)
	}

	statement, err = database.Prepare(
		"INSERT INTO people (firstname, lastname) VALUES (?, ?)",
	)
	if err != nil {
		log.Fatalf("sql.Prepare().error: %v", err)
	}

	_, err = statement.Exec("Rob", "Gronkowski")
	if err != nil {
		log.Fatalf("sql.Exec().error: %v", err)
	}

	rows, err = database.Query("SELECT id, firstname, lastname FROM people")
	if err != nil {
		log.Fatalf("sql.Query().error: %v", err)
	}

	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		err = rows.Scan(&id, &firstname, &lastname)
		if err != nil {
			log.Fatalf("rows.Scan().error: %v", err)
		}

		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}

	os.Exit(0)
	//var t ToJsObject = ToJsObject{
	//	"key": ToJsObject{
	//		"width": "100px",
	//	},
	//	"key_two": 1,
	//}
	//t.Convert()

	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static"))
	r.GET("/local/file", func(c *gin.Context) {
		c.File("local/file.go")
	})
	r.GET("/saveTimeLine", saveTimeLine)

	log.Println("Listening on :3000...")
	err = r.Run(":3000")
	if err != nil {
		log.Panic(err)
	}

}

func saveTimeLine(c *gin.Context) {
	var err error
	var body []byte

	body, err = ioutil.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", body)
}
