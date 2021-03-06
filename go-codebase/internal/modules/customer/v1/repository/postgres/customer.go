package postgres

import (
	"context"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gitlab.com/Wuriyanto/go-codebase/internal/modules/customer/v1/domain"
	"gitlab.com/Wuriyanto/go-codebase/pkg/shared"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CustomerRepositoryPostgres repository
type CustomerRepositoryPostgres struct {
	readDB, writeDB *gorm.DB
}

// NewCustomerRepositoryPostgres function will init CustomerRepositoryPostgres instance
func NewCustomerRepositoryPostgres(readDB, writeDB *gorm.DB) *CustomerRepositoryPostgres {
	return &CustomerRepositoryPostgres{readDB: readDB, writeDB: writeDB}
}

// Find function
func (r *CustomerRepositoryPostgres) Find(ctx context.Context, where domain.Customer) shared.Result {
	var customer domain.Customer

	query := make(map[string]interface{})
	query["is_deleted"] = false

	err := r.readDB.Where(&where, query).First(&customer).Error
	if err != nil {
		return shared.Result{Error: err}
	}

	return shared.Result{Data: &customer}
}

// FindAll function
func (r *CustomerRepositoryPostgres) FindAll(ctx context.Context) shared.Result {
	var filter domain.CustomerFilter
	filter.ParseFromContext(ctx)

	db := r.readDB

	if filter.OrderBy != "" {
		var sort = "desc"
		if filter.Sort == "asc" {
			sort = "asc"
		}
		db = db.Order(fmt.Sprintf("%s %s", filter.OrderBy, sort))
	}

	var Customers []domain.Customer
	err := db.Offset(filter.Offset).Limit(filter.Limit).Find(&Customers).Error
	if err != nil {
		return shared.Result{Error: err}
	}

	return shared.Result{Data: Customers}
}

// Count function
func (r *CustomerRepositoryPostgres) Count(ctx context.Context) shared.Result {
	var filter domain.CustomerFilter
	filter.ParseFromContext(ctx)

	db := r.readDB

	if filter.OrderBy != "" {
		var sort = "desc"
		if filter.Sort == "asc" {
			sort = "asc"
		}
		db = db.Order(fmt.Sprintf("%s %s", filter.OrderBy, sort))
	}

	var count int64
	err := db.Model(&domain.Customer{}).Count(&count).Error
	if err != nil {
		return shared.Result{Error: err}
	}

	return shared.Result{Data: count}
}

// Save function
func (r *CustomerRepositoryPostgres) Save(ctx context.Context, data *domain.Customer) shared.Result {
	err := r.writeDB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(data).Error

	if err != nil {
		return shared.Result{Error: err}
	}
	return shared.Result{Data: data}
}

// SaveBatch function
func (r *CustomerRepositoryPostgres) SaveBatch(ctx context.Context, datas []*domain.Customer) shared.Result {
	err := r.writeDB.Debug().Create(&datas).Error
	if err != nil {
		return shared.Result{Error: err}
	}
	return shared.Result{Data: datas}
}

// Update function
func (r *CustomerRepositoryPostgres) Update(ctx context.Context, data *domain.Customer) shared.Result {
	return shared.Result{}
}

// FindOne
func (r *CustomerRepositoryPostgres) FindOne(ctx context.Context, id uuid.UUID) shared.Result {
	var customer domain.Customer
	err := r.readDB.First(&customer, "id = ?", id.String()).Error
	if err != nil {
		return shared.Result{Error: err}
	}

	return shared.Result{Data: &customer}
}
