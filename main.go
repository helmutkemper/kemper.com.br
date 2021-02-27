package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
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
	var uniqueKeyBase = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "!", "@", "#", "$", "%", "&", "*", "(", ")", "_", "=", "-", "+", "[", "]", "{", "}", "|", "/", "?", "<", ">", ";", ":"}
	var randonGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i != 50; i += 1 {
		uniqueKey += uniqueKeyBase[randonGenerator.Intn(82-1)]
	}

	return
}

type Password struct{}

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

func (e *BusinessRules) Login(mail, password string) (successful bool, err error) {
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
