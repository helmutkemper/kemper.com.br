package server

import (
	"github.com/gin-gonic/gin"
	endpointLang "github.com/helmutkemper/kemper.com.br/gin/endpoint/languages"
	endpointMenu "github.com/helmutkemper/kemper.com.br/gin/endpoint/menu"
	endpointUser "github.com/helmutkemper/kemper.com.br/gin/endpoint/user"
	"log"
	"net/http"
)

func ConfigAndStart() (err error) {
	var epMenu = endpointMenu.DataSource{}
	var epUser = endpointUser.DataSource{}
	var epLanguages = endpointLang.DataSource{}

	r := gin.Default()
	r.StaticFS("/static", http.Dir("./cmd/static"))
	r.GET("/local/file", func(c *gin.Context) {
		c.File("local/file.go")
	})
	r.GET("/datasource/menu", epMenu.MenuMain)
	r.GET("/datasource/menuClassRoom", epMenu.MenuGetListClassRoom)
	r.GET("/datasource/user/:mail", epUser.UserByEmail)
	r.GET("/datasource/languages", epLanguages.GetAll)

	log.Println("Listening on :3000...")
	err = r.Run(":3000")
	return
}
