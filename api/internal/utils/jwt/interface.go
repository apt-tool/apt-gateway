package jwt

type Authenticator interface {
	GenerateToken(email string) (string, error)
	ParseToken(token string) (string, error)
}

func New(cfg Config) Authenticator {
	return &authenticator{
		key:    cfg.PrivateKey,
		expire: cfg.ExpireTime,
	}
}
