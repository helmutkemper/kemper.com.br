package businessRules

import "errors"

func (e *BusinessRules) UserLogin(mail, password string) (successful bool, err error) {
	var hashPassword string

	successful, err = e.MailSyntax(mail)
	if err != nil {
		return
	}
	if successful == false {
		err = errors.New("e-mail must be a valid syntax")
		return
	}

	successful, err = e.Data.MailExists(mail)
	if err != nil {
		return
	}
	if successful == false {
		return false, errors.New("email not found")
	}

	successful, err = e.Data.MailHasVerified(mail)
	if err != nil {
		return
	}
	if successful == false {
		return false, errors.New("email has not yet been verified")
	}

	hashPassword, err = e.Data.GetPassword(mail)
	if err != nil {
		return
	}

	successful = e.Password.CheckHash([]byte(password), []byte(hashPassword))
	if successful != true {
		return false, errors.New("password does not match")
	}
	return
}
