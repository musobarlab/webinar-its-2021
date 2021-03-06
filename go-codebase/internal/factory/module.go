package factory

import (
	"gitlab.com/Wuriyanto/go-codebase/internal/factory/base"
	"gitlab.com/Wuriyanto/go-codebase/internal/factory/interfaces"
	"gitlab.com/Wuriyanto/go-codebase/internal/modules/auth"
	"gitlab.com/Wuriyanto/go-codebase/internal/modules/customer"
)

// ModuleFactory factory
type ModuleFactory interface {
	RestHandler(version string) interfaces.EchoRestDelivery
	GRPCHandler() interfaces.GRPCDelivery
	GraphQLHandler() (name string, resolver interface{})
	SubscriberHandler(subsType base.Subscriber) interfaces.SubscriberDelivery
	Name() base.Module
}

// InitAllModule in this app
func InitAllModule(params *base.ModuleParam) []ModuleFactory {
	modules := []ModuleFactory{
		customer.NewModule(params),
		auth.NewModule(params),
	}

	return modules
}
