package auth

import (
	"gitlab.com/Wuriyanto/go-codebase/internal/factory/base"
	"gitlab.com/Wuriyanto/go-codebase/internal/factory/interfaces"
	authDeliveryV1 "gitlab.com/Wuriyanto/go-codebase/internal/modules/auth/v1/delivery"
	delivery "gitlab.com/Wuriyanto/go-codebase/internal/modules/auth/v1/delivery"
	authRepositoryV1 "gitlab.com/Wuriyanto/go-codebase/internal/modules/auth/v1/repository"
	authUsecaseV1 "gitlab.com/Wuriyanto/go-codebase/internal/modules/auth/v1/usecase"
	"gitlab.com/Wuriyanto/go-codebase/pkg/helper"
)

// Module model
type Module struct {
	V1 struct {
		restHandler *delivery.RestAuthHandler
	}
}

// NewModule module constructor
func NewModule(params *base.ModuleParam) *Module {
	repo := authRepositoryV1.NewRepository(params.Config)
	uc := authUsecaseV1.NewAuthUsecase(repo, params.Publisher)
	restV1Handler := authDeliveryV1.NewRestAuthHandler(uc, params.Middleware)

	var module Module
	module.V1.restHandler = restV1Handler
	return &module
}

// RestHandler method
func (m *Module) RestHandler(version string) (d interfaces.EchoRestDelivery) {
	switch version {
	case helper.V1:
		d = m.V1.restHandler
	case helper.V2:
		d = nil
	}
	return
}

// GRPCHandler method
func (m *Module) GRPCHandler() interfaces.GRPCDelivery {
	return nil
}

// GraphQLHandler method
func (m *Module) GraphQLHandler() (name string, resolver interface{}) {
	return "Auth", nil
}

// SubscriberHandler method
func (m *Module) SubscriberHandler(subsType base.Subscriber) (s interfaces.SubscriberDelivery) {
	return s
}

// Name get module name
func (m *Module) Name() base.Module {
	return base.ModuleAuth
}
