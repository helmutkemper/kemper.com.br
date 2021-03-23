package server

import (
	"github.com/gin-gonic/gin"
	endpointMenu "github.com/helmutkemper/kemper.com.br/gin/endpoint/menu"
	endpointUser "github.com/helmutkemper/kemper.com.br/gin/endpoint/user"
	"log"
	"net/http"
)

func ConfigAndStart() (err error) {
	var epMenu = endpointMenu.DataSource{}
	var epUser = endpointUser.DataSource{}

	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static"))
	r.GET("/local/file", func(c *gin.Context) {
		c.File("local/file.go")
	})
	r.GET("/datasource/menu", epMenu.Menu)
	r.GET("/datasource/user/:mail", epUser.Menu)

	log.Println("Listening on :3000...")
	err = r.Run(":3000")
	return
}
