package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gitlab.com/Wuriyanto/go-codebase/pkg/helper"
	"gitlab.com/Wuriyanto/go-codebase/pkg/wrapper"
)

// ValidateBearer jwt token middleware
func (m *Middleware) ValidateBearer() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authorization := c.Request().Header.Get(echo.HeaderAuthorization)
			if authorization == "" {
				return wrapper.NewHTTPResponse(http.StatusUnauthorized, "Invalid authorization").JSON(c.Response())
			}

			authValues := strings.Split(authorization, " ")
			authType := strings.ToLower(authValues[0])
			if authType != "bearer" || len(authValues) != 2 {
				return wrapper.NewHTTPResponse(http.StatusUnauthorized, "Invalid authorization").JSON(c.Response())
			}

			tokenString := authValues[1]
			resp := <-m.tokenUtils.Validate(c.Request().Context(), tokenString)
			if resp.Error != nil {
				return wrapper.NewHTTPResponse(http.StatusUnauthorized, resp.Error.Error()).JSON(c.Response())
			}

			c.Set(helper.TokenClaimKey, resp.Data)
			return next(c)
		}
	}
}
