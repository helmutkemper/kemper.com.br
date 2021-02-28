package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	sqlite3 "github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/dataAccess/SQLite3"
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/dataAccess/dataFormat"
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/view/dataSource"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	var err error

	var menu []dataFormat.Menu
	var sql3 = sqlite3.SQLite3{}
	err = sql3.Connect("./db/database.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	menu, err = sql3.GetMenu(0, nil)
	if err != nil {
		log.Fatal(err)
	}

	var dm = dataSource.Menu{}
	dm.CommonDataConvert(&menu, nil)

	var b []byte
	b, err = json.Marshal(&dm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n\n\n%s\n\n\n", b)

	os.Exit(0)

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
