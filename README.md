# kemper.com.br
GOROOT=/Users/kemper/go/go1.16 #gosetup
GOPATH=/Users/kemper/go/Libraries:/Users/kemper/go/projetos/kemper.com.br #gosetup
/Users/kemper/go/go1.16/bin/go test -c -o /private/var/folders/85/c_g4_d35425d5039h6tm3y040000gn/T/___ExampleSQLiteMenu_GetMenu_in_github_com_helmutkemper_kemper_com_br_src_github_com_helmutkemper_dataAccess_SQLite3 github.com/helmutkemper/kemper.com.br/src/github.com/helmutkemper/dataAccess/SQLite3 #gosetup
/Users/kemper/go/go1.16/bin/go tool test2json -t /private/var/folders/85/c_g4_d35425d5039h6tm3y040000gn/T/___ExampleSQLiteMenu_GetMenu_in_github_com_helmutkemper_kemper_com_br_src_github_com_helmutkemper_dataAccess_SQLite3 -test.v -test.paniconexit0 -test.run ^\QExampleSQLiteMenu_GetMenu\E$
=== RUN   ExampleSQLiteMenu_GetMenu
--- FAIL: ExampleSQLiteMenu_GetMenu (0.00s)
got:
[{Id:1 IdMenu:1 IdSecondary:0 Text:Kemper.com.br Admin:0 Icon:fas fa-code-branch Url: ItemOrder:0 Menu:[]} {Id:2 IdMenu:1 IdSecondary:0 Text:Códigos Admin:0 Icon:fas fa-code Url: ItemOrder:1 Menu:[]} {Id:3 IdMenu:1 IdSecondary:0 Text:Admin Admin:1 Icon:fas fa-cogs Url: ItemOrder:2 Menu:[]} {Id:4 IdMenu:1 IdSecondary:0 Text:Donation Admin:0 Icon: Url: ItemOrder:3 Menu:[]}]
want:


FAIL

Process finished with exit code 1

