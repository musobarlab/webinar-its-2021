package usecase

import (
	"context"
	"errors"

	uuid "github.com/satori/go.uuid"

	"gitlab.com/Wuriyanto/go-codebase/config"
	"gitlab.com/Wuriyanto/go-codebase/internal/modules/auth/v1/domain"
	"gitlab.com/Wuriyanto/go-codebase/internal/modules/auth/v1/repository"

	custDomain "gitlab.com/Wuriyanto/go-codebase/internal/modules/customer/v1/domain"

	"gitlab.com/Wuriyanto/go-codebase/pkg/messaging"
)

type authUsecaseImpl struct {
	repo      *repository.Repository
	publisher messaging.Publisher
}

// NewAuthUsecase constructor
func NewAuthUsecase(repo *repository.Repository, publisher messaging.Publisher) AuthUsecase {
	return &authUsecaseImpl{
		repo:      repo,
		publisher: publisher,
	}
}

func (uc *authUsecaseImpl) Login(ctx context.Context, request *domain.RequestLogin) (resp *domain.ResponseLogin, err error) {

	repoRes := uc.repo.Customer.Find(ctx, custDomain.Customer{
		Email: request.Username,
	})

	if repoRes.Error != nil {
		return nil, errors.New("invalid username or password")
	}

	customer := repoRes.Data.(*custDomain.Customer)

	if !customer.IsValidPassword(request.Password) {
		return nil, errors.New("invalid username or password")
	}

	// Generate access token & refresh token
	var tokenClaim domain.TokenClaim
	tokenClaim.User.Email = customer.Email
	tokenClaim.User.ID = customer.ID.String()
	tokenClaim.Alg = domain.RS256
	tokenClaim.JTI = uuid.NewV4().String()
	tokenClaim.RefreshJTI = uuid.NewV4().String()
	repoRes = <-uc.repo.Token.Generate(ctx, &tokenClaim, config.GlobalEnv.AccessTokenExpired)
	if repoRes.Error != nil {
		return nil, repoRes.Error
	}
	token := repoRes.Data.(string)

	var refreshTokenClaim domain.TokenClaim
	refreshTokenClaim.User.ID = customer.ID.String()
	refreshTokenClaim.Alg = domain.HS256
	refreshTokenClaim.JTI = tokenClaim.RefreshJTI
	repoRes = <-uc.repo.Token.Generate(ctx, &refreshTokenClaim, config.GlobalEnv.RefreshTokenExpired)
	if repoRes.Error != nil {
		return nil, repoRes.Error
	}
	refreshToken := repoRes.Data.(string)

	return &domain.ResponseLogin{
		AccessToken:           token,
		Role:                  "",
		RefreshToken:          refreshToken,
		AccessTokenExpiresIn:  int(config.GlobalEnv.AccessTokenExpired.Seconds()),
		RefreshTokenExpiresIn: int(config.GlobalEnv.RefreshTokenExpired.Seconds()),
	}, nil
}

func (uc *authUsecaseImpl) RefreshToken(ctx context.Context, refreshToken string) (*domain.RefreshTokenResponse, error) {
	return nil, nil
}

func (uc *authUsecaseImpl) Logout(ctx context.Context) (*domain.ResponseLogout, error) {
	return nil, nil
}
