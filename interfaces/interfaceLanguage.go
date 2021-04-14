package interfaces

type InterfaceLanguage interface {
	Connect(connectionString string, args ...interface{}) (err error)
	Close() (err error)
	Install() (err error)
	GetAll() (languagues []string, err error)
	Set(id, name string) (err error)
}
