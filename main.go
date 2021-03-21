package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/helmutkemper/kemper.com.br/businessRules/system/datasource"
	"github.com/helmutkemper/kemper.com.br/gin/server"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
)

func initSystemRules() (err error) {
	err = datasource.Linker.Init(datasource.KSQLite)
	return
}

func main() {
	var err error

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err = initSystemRules()
	if err != nil {
		log.Fatalf("System datasource initialization error: %v", err.Error())
	}

	err = server.ConfigAndStart()
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
