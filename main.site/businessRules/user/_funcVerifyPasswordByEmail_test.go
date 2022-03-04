package user

import (
	"github.com/helmutkemper/kemper.com.br/businessRules/system/datasource"
	dsUser "github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLite/_SQLiteUser"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/consts"
	"github.com/helmutkemper/kemper.com.br/toModule/passwordHash"
	"log"
	"os"
)

func ExampleBusinessRules_VerifyPasswordByEmail() {
	var err error
	var matched bool

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
	rules.Password = &passwordHash.Password{}
	err = rules.DataSource.Connect(consts.KDatabaseName)
	if err != nil {
		log.Fatalf("ds.Connect().error: %v", err.Error())
	}

	err = rules.Set("", 1, "admin", "admin", "admin@admin.com", "@dminAdminAdmin")
	if err != nil {
		log.Fatalf("rules.SetMainMenu().error: %v", err.Error())
	}

	matched, err = rules.VerifyPasswordByEmail("admin@admin.com", "@dminAdminAdmin")
	if err != nil {
		log.Fatalf("rules.VerifyPasswordByEmail().error: %v", err.Error())
	}

	if matched == false {
		log.Fatalf("matched error: %v", matched)
	}

	err = os.Remove(consts.KDatabaseName)
	if err != nil {
		log.Fatalf("os.Remove().error: %v", err.Error())
	}

	// Output:
	//
}
