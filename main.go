package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/helmutkemper/kemper.com.br/gin/endpoint"
	"github.com/helmutkemper/kemper.com.br/rules/system/datasource"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"net/http"
)

func initSystemRules() (err error) {
	err = datasource.ReferencesList.Init(datasource.KSQLite)
	return
}

func main() {
	var err error

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err = initSystemRules()
	if err != nil {
		log.Fatalf("System datasource initialization error: %v", err.Error())
	}

	var ginEndpoint = endpoint.MenuDataSource{}

	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static"))
	r.GET("/local/file", func(c *gin.Context) {
		c.File("local/file.go")
	})
	r.GET("/saveTimeLine", saveTimeLine)
	r.GET("/datasource/menu", ginEndpoint.Menu)

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
