package user

import (
	"github.com/gin-gonic/gin"
	"github.com/helmutkemper/kemper.com.br/businessRules/user"
	"github.com/helmutkemper/kemper.com.br/gin/endpoint/restful"
	"github.com/helmutkemper/kemper.com.br/interfaces"
	"github.com/helmutkemper/kemper.com.br/view/viewUser"
)

type DataSource struct {
	restful.Restful
	DataSource interfaces.InterfaceUser `json:"-"`
}

// Menu: (Português): Endpoint menu para o datasource do componente Kendo UI JQuery Menu
func (e *DataSource) Menu(c *gin.Context) {
	var err error
	var userData viewUser.User
	var length int

	var email = c.Param("mail")

	menuBusinessRules := user.BusinessRules{}
	length, userData, err = menuBusinessRules.GetByEmail(email)

	e.Meta.Error = []string{}
	if err != nil {
		e.Meta.Success = false
		e.Meta.Error = []string{err.Error()}
		e.Object = []int{}
		c.JSON(200, e)
		return
	}

	e.Meta.Total = length
	e.Meta.Actual = length
	e.Meta.Success = true
	e.Object = []viewUser.User{
		userData,
	}

	c.JSON(200, e)
}
