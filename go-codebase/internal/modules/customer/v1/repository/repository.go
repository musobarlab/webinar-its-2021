package repository

import (
	"gitlab.com/Wuriyanto/go-codebase/internal/modules/customer/v1/repository/interfaces"
	"gitlab.com/Wuriyanto/go-codebase/internal/modules/customer/v1/repository/postgres"
	"gorm.io/gorm"
)

// Repository model
type Repository struct {
	CustomerRepositoryPostgres interfaces.CustomerRepository
}

// NewRepository constructor
func NewRepository(readDB, writeDB *gorm.DB) *Repository {
	return &Repository{
		CustomerRepositoryPostgres: postgres.NewCustomerRepositoryPostgres(readDB, writeDB),
	}
}
