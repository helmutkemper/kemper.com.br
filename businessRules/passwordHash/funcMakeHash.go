package passwordHash

import "golang.org/x/crypto/bcrypt"

func (e *Password) MakeHash(password []byte) (hash []byte, err error) {
	hash, err = bcrypt.GenerateFromPassword(password, 30)
	return
}
