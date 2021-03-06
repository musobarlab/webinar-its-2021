package usecase

import (
	"context"

	"gitlab.com/Wuriyanto/go-codebase/internal/modules/auth/v1/domain"
)

// AuthUsecase abstraction
type AuthUsecase interface {
	Login(ctx context.Context, request *domain.RequestLogin) (*domain.ResponseLogin, error)
	RefreshToken(ctx context.Context, refreshToken string) (*domain.RefreshTokenResponse, error)
	Logout(ctx context.Context) (*domain.ResponseLogout, error)
}
