package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	echo "github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gitlab.com/Wuriyanto/go-codebase/pkg/shared"
	"gitlab.com/Wuriyanto/go-codebase/pkg/token/mocks"
)

func TestMiddleware_ValidateBearer(t *testing.T) {

	var generateRepoResult = func(data shared.Result) <-chan shared.Result {
		output := make(chan shared.Result)
		go func() {
			defer close(output)
			output <- data
		}()
		return output
	}

	t.Run("Testcase #1: Positive, valid auth", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer validtoken")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler := echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, c.String(http.StatusOK, "hello"))
		})

		tokenUtilMock := new(mocks.Token)
		tokenUtilMock.On("Validate", mock.Anything, "validtoken").Return(generateRepoResult(shared.Result{Error: nil}))

		midd := &Middleware{
			tokenUtils: tokenUtilMock,
		}

		mw := midd.ValidateBearer()(handler)
		err := mw(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Code, http.StatusOK)
	})

	t.Run("Testcase #2: Negative, empty authorization header", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler := echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, c.String(http.StatusOK, "hello"))
		})

		tokenUtilMock := new(mocks.Token)
		tokenUtilMock.On("Validate", mock.Anything, mock.Anything).Return(generateRepoResult(shared.Result{Error: nil}))

		midd := &Middleware{
			tokenUtils: tokenUtilMock,
		}

		mw := midd.ValidateBearer()(handler)
		err := mw(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Code, http.StatusUnauthorized)
		assert.Contains(t, rec.Body.String(), "Invalid authorization")
	})

	t.Run("Testcase #3: Negative, invalid authorization type", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Basic 123")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler := echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, c.String(http.StatusOK, "hello"))
		})

		tokenUtilMock := new(mocks.Token)
		tokenUtilMock.On("Validate", mock.Anything, mock.Anything).Return(generateRepoResult(shared.Result{Error: nil}))

		midd := &Middleware{
			tokenUtils: tokenUtilMock,
		}

		mw := midd.ValidateBearer()(handler)
		err := mw(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Code, http.StatusUnauthorized)
		assert.Contains(t, rec.Body.String(), "Invalid authorization")
	})

	t.Run("Testcase #4: Negative, token is expired", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer test")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler := echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, c.String(http.StatusOK, "hello"))
		})

		tokenUtilMock := new(mocks.Token)
		tokenUtilMock.On("Validate", mock.Anything, mock.Anything).Return(generateRepoResult(shared.Result{Error: errors.New("Token is expired")}))

		midd := &Middleware{
			tokenUtils: tokenUtilMock,
		}

		mw := midd.ValidateBearer()(handler)
		err := mw(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Code, http.StatusUnauthorized)
		assert.Contains(t, rec.Body.String(), "Token is expired")
	})
}
