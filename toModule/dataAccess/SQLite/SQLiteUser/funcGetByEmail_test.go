package SQLiteUser

import (
	"encoding/json"
	"fmt"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/consts"
	"log"
	"os"
)

func ExampleSQLiteUser_GetByEmail() {
	var err error
	var user dataFormat.User
	var userAsByte []byte

	var sqlUser = SQLiteUser{}
	err = sqlUser.Connect(consts.KDatabaseName)
	if err != nil {
		log.Fatalf("sqlUser.Connect().error: %v", err.Error())
	}

	err = sqlUser.Install()
	if err != nil {
		log.Fatalf("sqlUser.Install().error: %v", err.Error())
	}

	var found bool
	found, err = sqlUser.MailExists("helmut.kemper@gmail.com")
	if err != nil {
		log.Fatalf("sqlUser.MailExists().error: %v", err.Error())
	}

	if found == false {
		log.Fatal("sqlUser.MailExists().found: false")
	}

	user, err = sqlUser.GetByEmail("helmut.kemper@gmail.com")
	if err != nil {
		log.Fatalf("sqlUser.GetMainMenu().error: %v", err.Error())
	}

	userAsByte, err = json.Marshal(&user)
	if err != nil {
		log.Fatalf("json.Marshal().error: %v", err.Error())
	}

	fmt.Printf("%s", userAsByte)

	err = os.Remove(consts.KDatabaseName)
	if err != nil {
		log.Fatalf("os.Remove().error: %v", err.Error())
	}

	// Output:
	// {"id":1,"manuId":1,"admin":1,"name":"Helmut Kemper","nickname":"Kemper","mail":"helmut.kemper@gmail.com"}
}
