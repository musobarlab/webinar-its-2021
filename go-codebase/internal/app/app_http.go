package app

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	ew "github.com/labstack/echo/v4/middleware"
	"gitlab.com/Wuriyanto/go-codebase/config"
	"gitlab.com/Wuriyanto/go-codebase/pkg/helper"
	"gitlab.com/Wuriyanto/go-codebase/pkg/wrapper"
)

// ServeHTTP user service
func (a *App) ServeHTTP() {

	a.httpServer.HTTPErrorHandler = wrapper.CustomHTTPErrorHandler
	a.httpServer.Use(ew.CORSWithConfig(ew.DefaultCORSConfig))
	a.httpServer.Use(ew.Logger())

	apiGroup := a.httpServer.Group("/api")
	a.httpServer.GET("/", func(c echo.Context) error {
		return c.String(200, "Service up and running")
	})

	v1Group := apiGroup.Group(helper.V1)
	v2Group := apiGroup.Group(helper.V2)
	for _, m := range a.modules {
		if h := m.RestHandler(helper.V1); h != nil {
			h.Mount(v1Group)
		}
		if h := m.RestHandler(helper.V2); h != nil {
			h.Mount(v2Group)
		}
	}

	if err := a.httpServer.Start(fmt.Sprintf(":%d", config.GlobalEnv.HTTPPort)); err != nil {
		log.Println(err)
	}
}
