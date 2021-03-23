package menu

import (
	"github.com/helmutkemper/kemper.com.br/interfaces"
)

type BusinessRules struct {
	DataSource interfaces.InterfaceMenu `json:"-"`
	UniqueID   interfaces.InterfaceUID  `json:"-"`
}
