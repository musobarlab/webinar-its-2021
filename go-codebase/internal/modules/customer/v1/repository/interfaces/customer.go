package interfaces

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"gitlab.com/Wuriyanto/go-codebase/internal/modules/customer/v1/domain"
	"gitlab.com/Wuriyanto/go-codebase/pkg/shared"
)

// CustomerRepository abstraction
type CustomerRepository interface {
	Find(ctx context.Context, where domain.Customer) shared.Result
	FindAll(ctx context.Context) shared.Result
	Count(ctx context.Context) shared.Result
	Save(ctx context.Context, data *domain.Customer) shared.Result
	SaveBatch(ctx context.Context, datas []*domain.Customer) shared.Result
	Update(ctx context.Context, data *domain.Customer) shared.Result
	FindOne(ctx context.Context, id uuid.UUID) shared.Result
}
