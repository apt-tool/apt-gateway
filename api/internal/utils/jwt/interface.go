package jwt

type Authenticator interface {
	GenerateToken(email string) (string, error)
	ParseToken(token string) (string, error)
}
