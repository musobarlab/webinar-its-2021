package shared

import "context"

// Repository abstraction
type Repository interface {
	Find(ctx context.Context, target interface{}) error
	Save(ctx context.Context, data interface{}) error
	Delete(ctx context.Context) error
}
