package passwordHash

import (
	"bytes"
	"errors"
)

func (e *Password) ruleOneSpecialChars(password []byte) (err error) {
	var char []byte
	var specialChars = [][]byte{[]byte("`"), []byte("~"), []byte("!"), []byte("@"), []byte("#"), []byte("$"),
		[]byte("%"), []byte("^"), []byte("&"), []byte("*"), []byte("("), []byte(")"), []byte("-"), []byte("_"),
		[]byte("+"), []byte("="), []byte("["), []byte("{"), []byte("]"), []byte("}"), []byte("|"), []byte("\\"),
		[]byte(";"), []byte(":"), []byte("\""), []byte("'"), []byte("<"), []byte(">"), []byte(","), []byte("."),
		[]byte("/"), []byte("?")}
	for _, char = range specialChars {
		if bytes.Contains(password, char) {
			return
		}
	}

	err = errors.New("the password must be one special char")
	return
}
