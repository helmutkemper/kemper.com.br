package SQLiteMenu

import (
	"encoding/json"
	"fmt"
	"github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/dataAccess/dataFormat"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func ExampleSQLiteMenu_Get() {
	var err error
	var menu []dataFormat.Menu

	sqlMenu := SQLiteMenu{}
	err = sqlMenu.Connect(kTestDatabaseName)
	if err != nil {
		log.Fatalf("sqlMenu.Connect().error: %v", err.Error())
	}

	err = sqlMenu.CreateTable()
	if err != nil {
		log.Fatalf("sqlMenu.CreateTable().error: %v", err.Error())
	}

	err = sqlMenu.populateMenuForTest()
	if err != nil {
		log.Fatalf("sqlMenu.populateMenuForTest().error: %v", err.Error())
	}

	menu, err = sqlMenu.Get(1)
	if err != nil {
		log.Fatalf("sqlMenu.Get().error: %v", err.Error())
	}

	b, _ := json.Marshal(&menu)

	fmt.Printf("%s", b)

	err = os.Remove(kTestDatabaseName)
	if err != nil {
		log.Fatalf("os.Remove().error: %v", err.Error())
	}

	// Output:
	// [{"id":1,"menuId":1,"secondaryId":0,"text":"Kemper.com.br","admin":0,"icon":"fas fa-code-branch","url":"","itemOrder":0,"menu":[{"id":5,"menuId":1,"secondaryId":1,"text":"Sobre mim","admin":0,"icon":"fas fa-info-circle","url":"","itemOrder":0,"menu":[]},{"id":6,"menuId":1,"secondaryId":1,"text":"Github","admin":0,"icon":"fas fa-info-circle","url":"https://github.com/helmutkemper","itemOrder":1,"menu":[]},{"id":7,"menuId":1,"secondaryId":1,"text":"LinkedIn","admin":0,"icon":"fab fa-linkedin","url":"https://www.linkedin.com/in/helmut-kemper-93a5441b/","itemOrder":2,"menu":[]}]},{"id":2,"menuId":1,"secondaryId":0,"text":"Códigos","admin":0,"icon":"fas fa-code","url":"","itemOrder":1,"menu":[{"id":8,"menuId":1,"secondaryId":2,"text":"Conversando com o sênior","admin":0,"icon":"fas fa-fire-extinguisher","url":"","itemOrder":0,"menu":[]},{"id":9,"menuId":1,"secondaryId":2,"text":"Migrando para o Go","admin":0,"icon":"fas fa-fire","url":"","itemOrder":1,"menu":[]}]},{"id":3,"menuId":1,"secondaryId":0,"text":"Admin","admin":1,"icon":"fas fa-cogs","url":"","itemOrder":2,"menu":[{"id":11,"menuId":1,"secondaryId":3,"text":"New content","admin":1,"icon":"","url":"","itemOrder":0,"menu":[]},{"id":10,"menuId":1,"secondaryId":3,"text":"Login","admin":1,"icon":"","url":"","itemOrder":1,"menu":[]}]},{"id":4,"menuId":1,"secondaryId":0,"text":"Donation","admin":0,"icon":"","url":"","itemOrder":3,"menu":[]}]
}
