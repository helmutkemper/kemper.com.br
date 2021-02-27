package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type ToJsObject map[string]interface{}

func (e *ToJsObject) Convert() (jsObject []byte) {
	element := reflect.ValueOf(e).Elem()
	//mapKeys := element.MapKeys()
	//typeField := element.Type()

	switch element.Kind() {
	case reflect.Map:
		fmt.Printf("string: %v\n", element.Type().String())
		fmt.Printf("name: %v\n", element.Type().Name())
		fmt.Printf("keys: %v\n", element.MapKeys())

		for _, i := range element.MapKeys() {
			field := i.Type().Key()
			typeField := i.Type().String()

			_ = field
			_ = typeField
		}

		//e := element.Elem()
		//for i := 0; i < e.NumField(); i += 1 {
		//	field := e.Field(i)
		//	typeField := e.Type().Field(i)
		//
		//	_ = field
		//	_ = typeField
		//}
	case reflect.Bool:
	case reflect.Float32:
	case reflect.Float64:
	case reflect.Int:
	case reflect.Slice:
	case reflect.String:
	case reflect.Struct:
	case reflect.Interface:
		fmt.Printf("name: %v\n", element.Type().Name())
		fmt.Printf("key: %v\n", element.Type().Key())
	}

	fmt.Printf("kind: %v\n", element.Kind())

	os.Exit(1)

	for i := 0; i < element.NumField(); i += 1 {
		field := element.Field(i)
		typeField := element.Type().Field(i)

		_ = field
		_ = typeField
	}

	return nil
}

var uniqueKeyBase = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "!", "@", "#", "$", "%", "&", "*", "(", ")", "_", "=", "-", "+", "[", "]", "{", "}", "|", "/", "?", "<", ">", ";", ":"}

func UniqueKeyMake() (uniqueKey string) {

	var randonGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i != 50; i += 1 {
		uniqueKey += uniqueKeyBase[randonGenerator.Intn(82-1)]
	}

	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 30)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func DatabaseCreateTables(database *sql.DB) (err error) {
	err = DatabaseUserCreate(database)
	if err != nil {
		return
	}

	return
}

func DatabaseUserCreate(
	database *sql.DB,
) (
	err error,
) {

	var statement *sql.Stmt
	statement, err = database.Prepare(`
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

func DatabaseUserInsertVerifyMail(
	database *sql.DB,
	mail string,
) (
	err error,
	found bool,
) {

	var rows *sql.Rows

	rows, err = database.Query(`SELECT id FROM user WHERE mail = "?"`, mail)
	if err != nil {
		return
	}

	found = rows.Next()
	return
}

func DatabaseUserLoginVerifyMail(
	database *sql.DB,
	mail string,
) (
	err error,
	mailHasVerified bool,
) {

	var rows *sql.Rows

	rows, err = database.Query(`SELECT id FROM user WHERE mail = "?" AND mailVerified = 1`, mail)
	if err != nil {
		return
	}

	mailHasVerified = rows.Next()
	return
}

func DatabaseUserInsert(
	database *sql.DB,
	name,
	nickname string,
	gender,
	userType int,
	mail,
	password string,
	mailVerified int,
) (
	err error,
) {

	var statement *sql.Stmt

	statement, err = database.Prepare(`
		INSERT INTO people (
			name,
			nickname,
			gender,
			userType,
			mail,
			password,
			mailVerified
		) 
		VALUES (
			?,
			?,
			?,
			?,
			?,
			?,
			?
		)`,
	)
	if err != nil {
		return
	}

	_, err = statement.Exec(
		name,
		nickname,
		gender,
		userType,
		mail,
		password,
		mailVerified,
	)
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
