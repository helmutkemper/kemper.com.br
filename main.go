package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/helmutkemper/kemper.com.br/businessRules/system/datasource"
	"github.com/helmutkemper/kemper.com.br/gin/server"
	"github.com/helmutkemper/kemper.com.br/util"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
)

func main() {
	var err error

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err = datasource.Linker.Init(datasource.KMongoDB)
	if err != nil {
		util.TraceToLog()
		panic(err)
	}

	err = server.ConfigAndStart()
	if err != nil {
		util.TraceToLog()
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
