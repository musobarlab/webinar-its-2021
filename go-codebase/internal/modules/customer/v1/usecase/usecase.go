package usecase

import (
	"context"

	"gitlab.com/Wuriyanto/go-codebase/internal/modules/customer/v1/domain"
	"gitlab.com/Wuriyanto/go-codebase/pkg/shared"
)

// CustomerUsecase abstraction
type CustomerUsecase interface {
	FindAll(ctx context.Context, filter *domain.CustomerFilter) ([]domain.CustomerResponse, *shared.Meta, error)
	Register(ctx context.Context, data *domain.CustomerRequest) shared.Result
	UpdateProfile(ctx context.Context, data *domain.CustomerRequest) shared.Result
	GetProfile(ctx context.Context, id string) shared.Result
}
