package domain

import (
	"context"

	"gitlab.com/Wuriyanto/go-codebase/pkg/shared"
)

// CustomerFilter for get data customer
type CustomerFilter struct {
	shared.BaseFilter
	Customer
}

type CustomerFilterContextKey struct{}

var activeCustomerFilterContextKey = CustomerFilterContextKey{}

// SetToContext CustomerFilter value to context
func (f *CustomerFilter) SetToContext(ctx context.Context) context.Context {
	f.CalculateOffset()
	return context.WithValue(ctx, activeCustomerFilterContextKey, f)
}

// ParseFromContext parse CustomerFilter data from context
func (f *CustomerFilter) ParseFromContext(ctx context.Context) {
	if val, ok := ctx.Value(activeCustomerFilterContextKey).(*CustomerFilter); ok {
		*f = *val
	}
}
