package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

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
