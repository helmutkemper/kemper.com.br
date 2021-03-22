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

func ExampleBusinessRules_VerifyPasswordByEmail() {
	var err error
	var length int
	var matched bool
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

	err = rules.Set(1, 1, "admin", "admin", "admin@admin.com", "@dminAdminAdmin")
	if err != nil {
		log.Fatalf("rules.Set().error: %v", err.Error())
	}

	matched, err = rules.VerifyPasswordByEmail("admin@admin.com", "@dminAdminAdmin")
	if err != nil {
		log.Fatalf("rules.VerifyPasswordByEmail().error: %v", err.Error())
	}

	if matched == false {
		log.Fatalf("matched error: %v", matched)
	}

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
