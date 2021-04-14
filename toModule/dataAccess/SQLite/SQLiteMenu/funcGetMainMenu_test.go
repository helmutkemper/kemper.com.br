package SQLiteMenu

import (
	"encoding/json"
	"fmt"
	"github.com/helmutkemper/kemper.com.br/dataAccess/dataFormat"
	"github.com/helmutkemper/kemper.com.br/toModule/dataAccess/consts"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func ExampleSQLiteMenu_GetMainMenu() {
	var err error
	var menu []dataFormat.Menu
	var menuAsByte []byte
	var length int

	var sqlMenu = SQLiteMenu{}
	err = sqlMenu.Connect(consts.KDatabaseName)
	if err != nil {
		log.Fatalf("sqlMenu.Connect().error: %v", err.Error())
	}

	err = sqlMenu.Install()
	if err != nil {
		log.Fatalf("sqlMenu.Install().error: %v", err.Error())
	}

	menu, length, err = sqlMenu.GetMainMenu("5996b891-9d3c-4038-af37-cb07f5f0f72d")
	if err != nil {
		log.Fatalf("sqlMenu.GetMainMenu().error: %v", err.Error())
	}

	if length != 4 {
		log.Fatalf("sqlMenu.GetMainMenu()..lenght.error: %v. Must be 4", length)
	}

	menuAsByte, err = json.Marshal(&menu)
	if err != nil {
		log.Fatalf("json.Marshal().error: %v", err.Error())
	}

	fmt.Printf("%s", menuAsByte)

	err = os.Remove(consts.KDatabaseName)
	if err != nil {
		log.Fatalf("os.Remove().error: %v", err.Error())
	}

	// Output:
	// [{"id":"04daa10f-1059-446a-a2b9-b48d536c8b23","menuId":"5996b891-9d3c-4038-af37-cb07f5f0f72d","secondaryId":"","text":"Kemper.com.br","admin":0,"icon":"fas fa-code-branch","itemOrder":0,"menu":[{"id":"d791c1a2-fc36-454a-a7f4-038149e30e4a","menuId":"5996b891-9d3c-4038-af37-cb07f5f0f72d","secondaryId":"04daa10f-1059-446a-a2b9-b48d536c8b23","text":"Sobre mim","admin":0,"icon":"fas fa-info-circle","itemOrder":0},{"id":"d40894ba-8015-402f-9c10-5a4836413b1d","menuId":"5996b891-9d3c-4038-af37-cb07f5f0f72d","secondaryId":"04daa10f-1059-446a-a2b9-b48d536c8b23","text":"Github","admin":0,"icon":"fab fa-github","url":"https://github.com/helmutkemper","itemOrder":1},{"id":"873da901-8285-4846-85df-59f16850429e","menuId":"5996b891-9d3c-4038-af37-cb07f5f0f72d","secondaryId":"04daa10f-1059-446a-a2b9-b48d536c8b23","text":"LinkedIn","admin":0,"icon":"fab fa-linkedin","url":"https://www.linkedin.com/in/helmut-kemper-93a5441b/","itemOrder":2}]},{"id":"0409921a-600b-49db-b8cc-d906e3837082","menuId":"5996b891-9d3c-4038-af37-cb07f5f0f72d","secondaryId":"","text":"Códigos","admin":0,"icon":"fas fa-code","itemOrder":1,"menu":[{"id":"207b66f1-eef9-4572-a57d-5152b511e37d","menuId":"5996b891-9d3c-4038-af37-cb07f5f0f72d","secondaryId":"0409921a-600b-49db-b8cc-d906e3837082","text":"Conversando com o sênior","admin":0,"icon":"fas fa-fire-extinguisher","itemOrder":0},{"id":"ae116cdf-1941-4ca7-9129-0bad2793ca42","menuId":"5996b891-9d3c-4038-af37-cb07f5f0f72d","secondaryId":"0409921a-600b-49db-b8cc-d906e3837082","text":"Migrando para o Go","admin":0,"icon":"fas fa-fire","itemOrder":1}]},{"id":"1726ce36-7338-4932-858b-6863700dbe51","menuId":"5996b891-9d3c-4038-af37-cb07f5f0f72d","secondaryId":"","text":"Admin","admin":1,"icon":"fas fa-cogs","itemOrder":2,"menu":[{"id":"24fda64b-0af2-4da6-9d95-38664614d3bd","menuId":"5996b891-9d3c-4038-af37-cb07f5f0f72d","secondaryId":"1726ce36-7338-4932-858b-6863700dbe51","text":"New content","admin":1,"itemOrder":0},{"id":"08c7e064-8225-4ee4-9cf4-19dcdd546c84","menuId":"5996b891-9d3c-4038-af37-cb07f5f0f72d","secondaryId":"1726ce36-7338-4932-858b-6863700dbe51","text":"Login","admin":1,"itemOrder":1}]},{"id":"d6667be4-6794-4c6f-9eb2-0d3cac1d7801","menuId":"5996b891-9d3c-4038-af37-cb07f5f0f72d","secondaryId":"","text":"Donation","admin":0,"icon":"fas fa-donate","itemOrder":3}]
}
