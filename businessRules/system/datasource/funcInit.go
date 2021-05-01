package datasource

import (
	"errors"
	constants "github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/kemper.com.br/interfaces"
	jwtverify "github.com/helmutkemper/kemper.com.br/toModule/JWT"
	"github.com/helmutkemper/kemper.com.br/toModule/passwordHash"
	"github.com/helmutkemper/kemper.com.br/toModule/uID"
	"github.com/helmutkemper/util"
	"plugin"
)

// Init (Português): Inicializa o datasource escolhido
//   name: tyme Name
//     KSQLite: Inicializa o banco de dados como sendo o SQLite
func (e *RefList) Init(name Name) (err error) {

	err = errors.New(constants.KErrorInicializeDataSourceFirst)

	// Inicializa o objeto Password
	e.Password = &passwordHash.Password{}

	// Inicializa o objeto UID
	e.UniqueID = &uID.UID{}

	// Inicializa o gerador/verificador de JWT
	e.Jwt = &jwtverify.JwtVerify{}
	err = e.Jwt.NewAlgorithm([]byte("colocar em constants")) //fixme
	if err != nil {
		util.TraceToLog()
		return
	}

	// Inicializa o banco de dados
	switch name {

	case KMongoDB:
		err = e.installMenuByPlugin("./plugin/menu.mongodb.so")
		if err != nil {
			util.TraceToLog()
			return
		}

		err = e.installUserByPlugin("./plugin/user.mongodb.so")
		if err != nil {
			util.TraceToLog()
			return
		}

		err = e.installLanguageByPlugin("./plugin/languages.mongodb.so")
		if err != nil {
			util.TraceToLog()
			return
		}

	case KSQLite:
		err = e.installMenuByPlugin("./plugin/menu.sqlite.so")
		if err != nil {
			util.TraceToLog()
			return
		}

		err = e.installUserByPlugin("./plugin/user.sqlite.so")
		if err != nil {
			util.TraceToLog()
			return
		}

		err = e.installLanguageByPlugin("./plugin/languages.sqlite.so")
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	return
}

func (e *RefList) installMenuByPlugin(pluginPlath string) (err error) {
	var ok bool
	var menu *plugin.Plugin
	var menuSymbol plugin.Symbol

	menu, err = plugin.Open(pluginPlath)
	if err != nil {
		util.TraceToLog()
		return
	}

	menuSymbol, err = menu.Lookup("Menu")
	if err != nil {
		util.TraceToLog()
		return
	}

	e.Menu, ok = menuSymbol.(interfaces.InterfaceMenu)
	if ok == false {
		err = errors.New("plugin menu conversion into interface menu has an error")
		util.TraceToLog()
		return
	}

	_, err = e.Menu.New()
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}

func (e *RefList) installUserByPlugin(pluginPlath string) (err error) {
	var ok bool
	var user *plugin.Plugin
	var userSymbol plugin.Symbol

	user, err = plugin.Open(pluginPlath)
	if err != nil {
		util.TraceToLog()
		return
	}

	userSymbol, err = user.Lookup("User")
	if err != nil {
		util.TraceToLog()
		return
	}

	e.User, ok = userSymbol.(interfaces.InterfaceUser)
	if ok == false {
		err = errors.New("plugin user conversion into interface user has an error")
		util.TraceToLog()
		return
	}

	_, err = e.User.New()
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}

func (e *RefList) installLanguageByPlugin(pluginPlath string) (err error) {
	var ok bool
	var language *plugin.Plugin
	var languaSymbol plugin.Symbol

	language, err = plugin.Open(pluginPlath)
	if err != nil {
		util.TraceToLog()
		return
	}

	languaSymbol, err = language.Lookup("Language")
	if err != nil {
		util.TraceToLog()
		return
	}

	e.Language, ok = languaSymbol.(interfaces.InterfaceLanguage)
	if ok == false {
		err = errors.New("plugin language conversion into interface language has an error")
		util.TraceToLog()
		return
	}

	_, err = e.Language.New()
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
