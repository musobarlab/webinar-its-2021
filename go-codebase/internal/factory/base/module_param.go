package base

import (
	"gitlab.com/Wuriyanto/go-codebase/config"
	"gitlab.com/Wuriyanto/go-codebase/pkg/messaging"
	"gitlab.com/Wuriyanto/go-codebase/pkg/middleware"
)

// ModuleParam base
type ModuleParam struct {
	Config     *config.Config
	Middleware *middleware.Middleware
	Publisher  messaging.Publisher
}

// Module is the type returned by a Classifier module
type Module int

const (
	// ModuleAuth module
	ModuleAuth Module = iota
	// ModuleCustomer module
	ModuleCustomer
	// ModuleRole module
	ModuleRole
	// ModuleUser module
	ModuleUser
)
