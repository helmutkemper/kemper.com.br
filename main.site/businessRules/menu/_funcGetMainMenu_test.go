package menu

import (
	"fmt"
	"github.com/helmutkemper/kemper.com.br/businessRules/system/datasource"
	consts "github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	dsMenu "github.com/helmutkemper/kemper.com.br/toModule/dataAccess/SQLite/_SQLiteMenu"
	"github.com/helmutkemper/kemper.com.br/view/menu/viewMainMenu"
	"log"
	"os"
)

func ExampleBusinessRules_GetMainMenu() {
	var err error
	var length int
	var menu viewMainMenu.Menu

	var ds = dsMenu.SQLiteMenu{}
	err = ds.Connect(constants.KDatabaseName)
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
	rules.DataSource = &dsMenu.SQLiteMenu{}

	err = rules.DataSource.Connect(consts.KDatabaseName)
	if err != nil {
		log.Fatalf("ds.Connect().error: %v", err.Error())
	}

	length, menu, err = rules.GetMainMenu(consts2.KMainMenuID)
	if err != nil {
		log.Fatalf("rules.GetMainMenu().error: %v", err.Error())
	}

	if length != 4 {
		log.Fatalf("length size error: %v", length)
	}

	fmt.Printf("%+v", menu)

	err = os.Remove(consts.KDatabaseName)
	if err != nil {
		log.Fatalf("os.Remove().error: %v", err.Error())
	}

	// Output:
	// [{Text:Kemper.com.br SpriteCssClass:menuWith fas fa-code-branch DataSpriteCssClassField:k-test-menu Items:[{Text:Sobre mim SpriteCssClass:menuWith fas fa-info-circle DataSpriteCssClassField:k-test-menu Items:[] Url:} {Text:Github SpriteCssClass:menuWith fab fa-github DataSpriteCssClassField:k-test-menu Items:[] Url:https://github.com/helmutkemper} {Text:LinkedIn SpriteCssClass:menuWith fab fa-linkedin DataSpriteCssClassField:k-test-menu Items:[] Url:https://www.linkedin.com/in/helmut-kemper-93a5441b/}] Url:} {Text:Códigos SpriteCssClass:menuWith fas fa-code DataSpriteCssClassField:k-test-menu Items:[{Text:Conversando com o sênior SpriteCssClass:menuWith fas fa-fire-extinguisher DataSpriteCssClassField:k-test-menu Items:[] Url:} {Text:Migrando para o Go SpriteCssClass:menuWith fas fa-fire DataSpriteCssClassField:k-test-menu Items:[] Url:}] Url:} {Text:Admin SpriteCssClass:menuWith fas fa-cogs DataSpriteCssClassField:k-test-menu Items:[{Text:New content SpriteCssClass: DataSpriteCssClassField:k-test-menu Items:[] Url:} {Text:Login SpriteCssClass: DataSpriteCssClassField:k-test-menu Items:[] Url:}] Url:} {Text:Donation SpriteCssClass:menuWith fas fa-donate DataSpriteCssClassField:k-test-menu Items:[] Url:}]

}
