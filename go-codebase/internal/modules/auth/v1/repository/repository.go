package repository

import (
	"gitlab.com/Wuriyanto/go-codebase/config"
	"gitlab.com/Wuriyanto/go-codebase/internal/modules/auth/v1/repository/interfaces"
	"gitlab.com/Wuriyanto/go-codebase/internal/modules/auth/v1/repository/jwt"
	customerrepo "gitlab.com/Wuriyanto/go-codebase/internal/modules/customer/v1/repository/interfaces"
	customerpostgres "gitlab.com/Wuriyanto/go-codebase/internal/modules/customer/v1/repository/postgres"
)

// Repository model
type Repository struct {
	Customer          customerrepo.CustomerRepository
	Token             interfaces.TokenRepository
	Cache             interfaces.CacheRepository
	RefreshTokenCache interfaces.CacheRepository
}

// NewRepository constructor
func NewRepository(cfg *config.Config) *Repository {
	return &Repository{
		Token:    jwt.NewJWTRepository(cfg.PublicKey, cfg.PrivateKey),
		Customer: customerpostgres.NewCustomerRepositoryPostgres(cfg.ReadDB, cfg.WriteDB),
	}
}
