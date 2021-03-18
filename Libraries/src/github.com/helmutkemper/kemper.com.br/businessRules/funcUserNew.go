package businessRules

import "errors"

func (e *BusinessRules) UserNew(name, nickname string, gender int, mail, password string) (err error) {
	var pass bool

	pass, err = e.MailSyntax(mail)
	if err != nil {
		return
	}
	if pass == false {
		err = errors.New("e-mail must be a valid syntax")
		return
	}

	pass, err = e.Data.MailExists(mail)
	if err != nil {
		return
	}
	if pass == false {
		return errors.New("email found")
	}

	err = e.Password.NewPasswordRule([]byte(password))
	if err != nil {
		return
	}

	err = e.Data.UserInsert(name, nickname, gender, 0, mail, password, 0)
	return
}
