package passwordHash

import (
	"bytes"
	"errors"
)

func (e *Password) ruleLowerCase(password []byte) (err error) {
	var char []byte
	var specialChars = [][]byte{[]byte("a"), []byte("b"), []byte("c"), []byte("d"), []byte("e"), []byte("f"),
		[]byte("g"), []byte("h"), []byte("i"), []byte("j"), []byte("k"), []byte("l"), []byte("m"), []byte("n"),
		[]byte("o"), []byte("p"), []byte("q"), []byte("r"), []byte("s"), []byte("t"), []byte("v"), []byte("w"),
		[]byte("x"), []byte("y"), []byte("z")}
	for _, char = range specialChars {
		if bytes.Contains(password, char) {
			return
		}
	}

	err = errors.New("the password must be one lower case char")
	return
}
