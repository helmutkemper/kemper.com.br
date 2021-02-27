package interfaces

type DataAccess interface {
	Init() (err error)
	Connect(filePath string) (err error)
	MailExists(mail string) (found bool, err error)
	MailHasVerified(mail string) (mailHasVerified bool, err error)
	UserInsert(name, nickname string, gender, userType int, mail, password string, mailVerified int) (err error)
	GetPassword(mail string) (password string, err error)
}
