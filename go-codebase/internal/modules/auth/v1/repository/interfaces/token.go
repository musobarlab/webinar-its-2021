package interfaces

import (
	"context"
	"time"

	"gitlab.com/Wuriyanto/go-codebase/internal/modules/auth/v1/domain"
	"gitlab.com/Wuriyanto/go-codebase/pkg/shared"
)

// TokenRepository abstraction
type TokenRepository interface {
	Generate(ctx context.Context, payload *domain.TokenClaim, expired time.Duration) <-chan shared.Result
	Refresh(ctx context.Context, token string) <-chan shared.Result
	Validate(ctx context.Context, token string) <-chan shared.Result
	Revoke(ctx context.Context, token string) <-chan shared.Result
}

// CacheRepository abstraction
type CacheRepository interface {
	shared.Repository
	GetTTL(ctx context.Context, key string) int
}
