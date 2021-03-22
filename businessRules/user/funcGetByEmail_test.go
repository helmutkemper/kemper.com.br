package user

import (
	"fmt"
	"github.com/helmutkemper/kemper.com.br/businessRules/system/datasource"
	dsUser "github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLite/SQLiteUser"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/consts"
	"github.com/helmutkemper/kemper.com.br/view/viewUser"
	"log"
	"os"
)

func ExampleBusinessRules_GetByEmail() {
	var err error
	var length int
	var user viewUser.User

	var ds = dsUser.SQLiteUser{}
	err = ds.Connect(consts.KDatabaseName)
	if err != nil {
		log.Fatalf("ds.Connect().error: %v", err.Error())
	}

	err = ds.Install()
	if err != nil {
		log.Fatalf("ds.Install().error: %v", err.Error())
	}

	err = datasource.Linker.Init(datasource.KSQLite)
	if err != nil {
		log.Fatalf("datasource.Linker.Init().error: %v", err.Error())
	}

	var rules = BusinessRules{}
	rules.DataSource = &dsUser.SQLiteUser{}
	err = rules.DataSource.Connect(consts.KDatabaseName)
	if err != nil {
		log.Fatalf("ds.Connect().error: %v", err.Error())
	}

	length, user, err = rules.GetByEmail("helmut.kemper@gmail.com")

	if length != 1 {
		log.Fatalf("length size error: %v", length)
	}

	fmt.Printf("%+v", user)

	err = os.Remove(consts.KDatabaseName)
	if err != nil {
		log.Fatalf("os.Remove().error: %v", err.Error())
	}

	// Output:
	// {Id:1 MenuId:1 Admin:1 Name:Helmut Kemper NickName:Kemper Mail:helmut.kemper@gmail.com Password:}
}
