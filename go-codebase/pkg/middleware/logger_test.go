package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	e := echo.New()

	t.Run("Testcase #1", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler := echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, c.String(http.StatusOK, "hello"))
		})

		mw := Logger(handler)
		err := mw(c)
		assert.NoError(t, err)
	})

	t.Run("Testcase #2", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderXRealIP, "127.0.0.1")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler := echo.HandlerFunc(func(c echo.Context) error {
			return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
		})

		mw := Logger(handler)
		err := mw(c)
		assert.NoError(t, err)
	})

	t.Run("Testcase #3", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderXForwardedFor, "local")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler := echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, c.String(http.StatusOK, "hello"))
		})

		mw := Logger(handler)
		err := mw(c)
		assert.NoError(t, err)
	})

}
