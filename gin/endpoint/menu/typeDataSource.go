package menu

import (
	"github.com/helmutkemper/kemper.com.br/gin/endpoint/restful"
	"github.com/helmutkemper/kemper.com.br/interfaces"
)

type DataSource struct {
	restful.Restful
	DataSource interfaces.InterfaceMenu `json:"-"`
}
