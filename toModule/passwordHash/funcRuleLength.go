package passwordHash

import (
	"errors"
	"log"
)

func (e *Password) ruleLength(password []byte) (err error) {
	if len(password) < 8 {
		err = errors.New("the password must be 8 letters or more")
		log.Printf("passwordHash.ruleLength().error: %v", err.Error())
	}

	return
}
