package usecase

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"gitlab.com/Wuriyanto/go-codebase/internal/modules/customer/v1/domain"
	"gitlab.com/Wuriyanto/go-codebase/internal/modules/customer/v1/repository"
	"gitlab.com/Wuriyanto/go-codebase/pkg/shared"
)

type customerUsecaseImpl struct {
	repo *repository.Repository
}

// NewCustomerUsecase create new customer usecase
func NewCustomerUsecase(repo *repository.Repository) CustomerUsecase {
	return &customerUsecaseImpl{
		repo: repo,
	}
}

func (uc *customerUsecaseImpl) FindAll(ctx context.Context, filter *domain.CustomerFilter) ([]domain.CustomerResponse, *shared.Meta, error) {
	ctx = filter.SetToContext(ctx)
	findAllRes := uc.repo.CustomerRepositoryPostgres.FindAll(ctx)
	if findAllRes.Error != nil {
		return nil, nil, findAllRes.Error
	}

	countRes := uc.repo.CustomerRepositoryPostgres.Count(ctx)
	if countRes.Error != nil {
		return nil, nil, countRes.Error
	}

	customers := findAllRes.Data.([]domain.Customer)
	count := countRes.Data.(int64)
	meta := shared.NewMeta(filter.Page, filter.Limit, int(count))

	var customerResponses []domain.CustomerResponse
	for _, customer := range customers {
		customerResponse := domain.CustomerResponseFromCustomer(&customer)
		customerResponses = append(customerResponses, customerResponse)
	}
	return customerResponses, meta, nil

}

func (uc *customerUsecaseImpl) Register(ctx context.Context, data *domain.CustomerRequest) shared.Result {
	customer, err := data.ToCustomerModel()

	if err != nil {
		return shared.Result{Error: err}
	}

	saveResult := uc.repo.CustomerRepositoryPostgres.Save(ctx, customer)
	if saveResult.Error != nil {
		return shared.Result{Error: saveResult.Error}
	}

	savedCustomer := saveResult.Data.(*domain.Customer)

	return shared.Result{Data: domain.CustomerResponseFromCustomer(savedCustomer)}
}

func (uc *customerUsecaseImpl) UpdateProfile(ctx context.Context, data *domain.CustomerRequest) shared.Result {
	customerUUID, err := uuid.FromString(data.CustomerID)
	if err != nil {
		return shared.Result{Error: err}
	}

	findOneResult := uc.repo.CustomerRepositoryPostgres.FindOne(ctx, customerUUID)
	if findOneResult.Error != nil {
		return shared.Result{Error: findOneResult.Error}
	}

	customer := findOneResult.Data.(*domain.Customer)

	customer.CollectRequest(data)

	saveResult := uc.repo.CustomerRepositoryPostgres.Save(ctx, customer)
	if saveResult.Error != nil {
		return shared.Result{Error: saveResult.Error}
	}

	savedCustomer := saveResult.Data.(*domain.Customer)

	return shared.Result{Data: domain.CustomerResponseFromCustomer(savedCustomer)}
}

func (uc *customerUsecaseImpl) GetProfile(ctx context.Context, id string) shared.Result {
	customerUUID, err := uuid.FromString(id)
	if err != nil {
		return shared.Result{Error: err}
	}

	findOneResult := uc.repo.CustomerRepositoryPostgres.FindOne(ctx, customerUUID)
	if findOneResult.Error != nil {
		return shared.Result{Error: findOneResult.Error}
	}

	customer := findOneResult.Data.(*domain.Customer)

	if customer.Rank == domain.CustomerRankGold {
		customer.Bonus = 400000
	}

	saveResult := uc.repo.CustomerRepositoryPostgres.Save(ctx, customer)
	if saveResult.Error != nil {
		return shared.Result{Error: saveResult.Error}
	}

	savedCustomer := saveResult.Data.(*domain.Customer)

	return shared.Result{Data: domain.CustomerResponseFromCustomer(savedCustomer)}
}
