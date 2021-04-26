package interfaces

type InterfaceJWT interface {
	NewAlgorithm(secretKey []byte) (err error)
	BuildToken(userUID, tokenUID string, audience []string) (token []byte, err error)
	Verify(token []byte) (tokenUID, userUI string, err error)
}
