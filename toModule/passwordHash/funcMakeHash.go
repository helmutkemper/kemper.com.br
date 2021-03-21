package passwordHash

import "golang.org/x/crypto/bcrypt"

func (e *Password) MakeHash(password []byte) (hash []byte, err error) {
	err = e.newPasswordRules(password)
	if err != nil {
		return
	}

	hash, err = bcrypt.GenerateFromPassword(password, 30)
	return
}
