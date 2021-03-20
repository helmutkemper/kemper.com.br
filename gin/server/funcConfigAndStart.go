package server

import (
	"github.com/gin-gonic/gin"
	"github.com/helmutkemper/kemper.com.br/gin/endpoint"
	"log"
	"net/http"
)

func ConfigAndStart() (err error) {
	var ginEndpoint = endpoint.MenuDataSource{}

	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static"))
	r.GET("/local/file", func(c *gin.Context) {
		c.File("local/file.go")
	})
	r.GET("/datasource/menu", ginEndpoint.Menu)

	log.Println("Listening on :3000...")
	err = r.Run(":3000")
	return
}
