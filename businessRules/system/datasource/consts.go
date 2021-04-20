package datasource

const (
	// KSQLite (Português): Define o datasource como sendo o SQLite3
	KSQLite Name = iota

	// KMongoDB (Português): Define o datasource como sendo o MongoDB
	KMongoDB

	// KTotalMadness (Português): Feito apenas para demonstração
	KTotalMadness

	// KPlugin (Português): Carrega um plugin externo com o módulo escolhido
	KPlugin
)
