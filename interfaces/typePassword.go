package interfaces

type Password interface {
	MakeHash(password []byte) (hash []byte, err error)
	CheckHash(password, hash []byte) (match bool)
	NewPasswordRule(password []byte) (err error)
}
