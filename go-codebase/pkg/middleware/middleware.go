package middleware

import (
	"gitlab.com/Wuriyanto/go-codebase/config"
	"gitlab.com/Wuriyanto/go-codebase/pkg/token"
)

// Middleware model
type Middleware struct {
	tokenUtils         token.Token
	username, password string
	grpcAuthKey        string
}

// NewMiddleware create new middleware instance
func NewMiddleware(cfg *config.Config) *Middleware {
	return &Middleware{
		tokenUtils:  token.NewJWT(cfg.PublicKey, cfg.PrivateKey),
		username:    config.GlobalEnv.BasicAuthUsername,
		password:    config.GlobalEnv.BasicAuthPassword,
		grpcAuthKey: config.GlobalEnv.GRPCAuthKey,
	}
}
