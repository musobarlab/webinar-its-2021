package app

import (
	"context"

	"github.com/labstack/echo/v4"
	"gitlab.com/Wuriyanto/go-codebase/api/jsonschema"
	"gitlab.com/Wuriyanto/go-codebase/config"
	"gitlab.com/Wuriyanto/go-codebase/internal/factory"
	"gitlab.com/Wuriyanto/go-codebase/internal/factory/base"
	"gitlab.com/Wuriyanto/go-codebase/pkg/middleware"
)

// App definition
type App struct {
	config     *config.Config
	modules    []factory.ModuleFactory
	httpServer *echo.Echo
}

// New  app
func New(cfg *config.Config) *App {
	jsonschema.Load(config.GlobalEnv.RootApp + "/api/jsonschema")

	mw := middleware.NewMiddleware(cfg)
	params := &base.ModuleParam{
		Config:     cfg,
		Middleware: mw,
	}
	modules := factory.InitAllModule(params)

	// init http server
	echoServer := echo.New()

	return &App{
		config:     cfg,
		modules:    modules,
		httpServer: echoServer,
	}
}

// Shutdown graceful shutdown all server, panic if there is still a process running when the request exceed given timeout in context
func (a *App) Shutdown(ctx context.Context) {
	if err := a.httpServer.Shutdown(ctx); err != nil {
		panic(err)
	}
}
