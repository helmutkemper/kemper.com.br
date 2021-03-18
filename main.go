package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/endpoint"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var err error

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	menu := endpoint.MenuDataSource{}

	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static"))
	r.GET("/local/file", func(c *gin.Context) {
		c.File("local/file.go")
	})
	r.GET("/saveTimeLine", saveTimeLine)
	r.GET("/datasource/menu", menu.Menu)

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
