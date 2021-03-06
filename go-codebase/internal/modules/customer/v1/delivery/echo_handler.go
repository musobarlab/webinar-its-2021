package delivery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/Wuriyanto/go-codebase/api/jsonschema"
	"gitlab.com/Wuriyanto/go-codebase/internal/modules/customer/v1/domain"
	"gitlab.com/Wuriyanto/go-codebase/internal/modules/customer/v1/usecase"
	"gitlab.com/Wuriyanto/go-codebase/pkg/helper"
	"gitlab.com/Wuriyanto/go-codebase/pkg/middleware"
	"gitlab.com/Wuriyanto/go-codebase/pkg/token"
	"gitlab.com/Wuriyanto/go-codebase/pkg/wrapper"
)

// RestCustomerHandler handler
type RestCustomerHandler struct {
	midd            *middleware.Middleware
	customerUsecase usecase.CustomerUsecase
}

// NewRestCustomerHandler create new rest handler
func NewRestCustomerHandler(midd *middleware.Middleware, customerUsecase usecase.CustomerUsecase) *RestCustomerHandler {
	return &RestCustomerHandler{
		midd:            midd,
		customerUsecase: customerUsecase,
	}
}

// Mount v1 handler (/v1)
func (h *RestCustomerHandler) Mount(root *echo.Group) {
	//customer := root.Group("/customer", h.midd.ValidateBearer())
	customer := root.Group("/customers")

	customer.GET("", h.getAllCustomer)
	customer.GET("/profile", h.getProfile, h.midd.ValidateBearer())
	customer.PUT("/profile", h.updateProfile, h.midd.ValidateBearer())
	customer.POST("", h.register)
}

// POST localhost:8080/api/v1/customer
func (h *RestCustomerHandler) updateProfile(c echo.Context) error {
	claim := c.Get(helper.TokenClaimKey).(*token.Claim)
	userIDStr := claim.User.ID

	body, _ := ioutil.ReadAll(c.Request().Body)
	if err := jsonschema.ValidateDocument("customer/update_profile", body); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "error validate payload", err).JSON(c.Response())
	}

	var payload domain.CustomerRequest
	json.Unmarshal(body, &payload)

	payload.CustomerID = userIDStr

	registerResult := h.customerUsecase.UpdateProfile(c.Request().Context(), &payload)
	if registerResult.Error != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "error validate payload", registerResult.Error).JSON(c.Response())
	}

	customerResponse := registerResult.Data.(domain.CustomerResponse)

	return wrapper.NewHTTPResponse(http.StatusCreated, "update succeed", customerResponse).JSON(c.Response())
}

// GET localhost:8080/api/v1/customer
func (h *RestCustomerHandler) getAllCustomer(c echo.Context) error {
	var filter domain.CustomerFilter
	if err := helper.ParseFromQueryParam(c.Request().URL.Query(), &filter); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "error parsing filter", err).JSON(c.Response())
	}

	if err := jsonschema.Validate("customer/get_all", filter); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "error validate filter", err).JSON(c.Response())
	}

	data, meta, err := h.customerUsecase.FindAll(c.Request().Context(), &filter)
	if err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "error get all users", err).JSON(c.Response())
	}
	return wrapper.NewHTTPResponse(http.StatusOK, "get all user succeed", data, meta).JSON(c.Response())
}

// POST localhost:8080/api/v1/customer
func (h *RestCustomerHandler) register(c echo.Context) error {

	body, _ := ioutil.ReadAll(c.Request().Body)
	if err := jsonschema.ValidateDocument("customer/customer", body); err != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "error validate payload", err).JSON(c.Response())
	}

	var payload domain.CustomerRequest
	json.Unmarshal(body, &payload)

	registerResult := h.customerUsecase.Register(c.Request().Context(), &payload)
	if registerResult.Error != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "error validate payload", registerResult.Error).JSON(c.Response())
	}

	customerResponse := registerResult.Data.(domain.CustomerResponse)

	return wrapper.NewHTTPResponse(http.StatusCreated, "register succeed", customerResponse).JSON(c.Response())
}

// POST localhost:8080/api/v1/customer
func (h *RestCustomerHandler) getProfile(c echo.Context) error {
	claim := c.Get(helper.TokenClaimKey).(*token.Claim)
	userIDStr := claim.User.ID

	getProfileResult := h.customerUsecase.GetProfile(c.Request().Context(), userIDStr)
	if getProfileResult.Error != nil {
		return wrapper.NewHTTPResponse(http.StatusBadRequest, "error validate payload", getProfileResult.Error).JSON(c.Response())
	}

	customerResponse := getProfileResult.Data.(domain.CustomerResponse)

	return wrapper.NewHTTPResponse(http.StatusCreated, "get profile succeed", customerResponse).JSON(c.Response())
}
