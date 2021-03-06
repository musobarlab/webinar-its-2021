package delivery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/Wuriyanto/go-codebase/api/jsonschema"
	"gitlab.com/Wuriyanto/go-codebase/internal/modules/auth/v1/domain"
	"gitlab.com/Wuriyanto/go-codebase/internal/modules/auth/v1/usecase"
	"gitlab.com/Wuriyanto/go-codebase/pkg/middleware"
	"gitlab.com/Wuriyanto/go-codebase/pkg/shared"
	"gitlab.com/Wuriyanto/go-codebase/pkg/wrapper"
)

// RestAuthHandler handler
type RestAuthHandler struct {
	authUsecase usecase.AuthUsecase
	mw          *middleware.Middleware
}

// NewRestAuthHandler create new rest handler
func NewRestAuthHandler(authUsecase usecase.AuthUsecase, mw *middleware.Middleware) *RestAuthHandler {
	return &RestAuthHandler{
		authUsecase: authUsecase,
		mw:          mw,
	}
}

// Mount v1 handler (/v1)
func (h *RestAuthHandler) Mount(root *echo.Group) {
	auth := root.Group("/auth")

	auth.POST("/login", h.login)
}

// GET localhost:8080/api/v1/auth/login
func (h *RestAuthHandler) login(c echo.Context) error {

	body, _ := ioutil.ReadAll(c.Request().Body)
	if err := jsonschema.ValidateDocument("auth/login", body); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed to validate request data", err).JSON(c.Response())
	}

	var payload domain.RequestLogin
	json.Unmarshal(body, &payload)

	payload.ApplicationType = c.Request().Header.Get("ApplicationType")
	ipAddress := c.RealIP()
	ctx := shared.SetToContext(c.Request().Context(), shared.ContextKey("ipAddress"), ipAddress)
	response, err := h.authUsecase.Login(ctx, &payload)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusUnauthorized, "Failed to login", err).JSON(c.Response())
	}

	return wrapper.NewHTTPResponse(http.StatusOK, "Your Request Has Been Processed", response).JSON(c.Response())
}
