package customer

import (
	"gitlab.com/Wuriyanto/go-codebase/internal/factory/base"
	"gitlab.com/Wuriyanto/go-codebase/internal/factory/interfaces"
	customerDeliveryV1 "gitlab.com/Wuriyanto/go-codebase/internal/modules/customer/v1/delivery"
	customerRepositoryV1 "gitlab.com/Wuriyanto/go-codebase/internal/modules/customer/v1/repository"
	customerUsecaseV1 "gitlab.com/Wuriyanto/go-codebase/internal/modules/customer/v1/usecase"
	"gitlab.com/Wuriyanto/go-codebase/pkg/helper"
)

// Module model
type Module struct {
	V1 struct {
		RestHandler *customerDeliveryV1.RestCustomerHandler
	}
}

// NewModule module constructor
func NewModule(params *base.ModuleParam) *Module {
	// customer modules version 1
	customerRepoV1 := customerRepositoryV1.NewRepository(params.Config.ReadDB, params.Config.WriteDB)
	customerUcV1 := customerUsecaseV1.NewCustomerUsecase(customerRepoV1)
	customerRestHandlerV1 := customerDeliveryV1.NewRestCustomerHandler(params.Middleware, customerUcV1)

	// customer modules version 2
	// ...

	// another modules version 1
	// ...

	var module Module
	module.V1.RestHandler = customerRestHandlerV1
	return &module
}

// RestHandler method
func (m *Module) RestHandler(version string) (d interfaces.EchoRestDelivery) {
	switch version {
	case helper.V1:
		d = m.V1.RestHandler
	case helper.V2:
		d = nil // TODO versioning
	}
	return
}

// GRPCHandler method
func (m *Module) GRPCHandler() interfaces.GRPCDelivery {
	return nil
}

// GraphQLHandler method
func (m *Module) GraphQLHandler() (name string, resolver interface{}) {
	return "Customer", nil
}

// SubscriberHandler method
func (m *Module) SubscriberHandler(subsType base.Subscriber) interfaces.SubscriberDelivery {
	return nil
}

// Name get module name
func (m *Module) Name() base.Module {
	return base.ModuleCustomer
}
