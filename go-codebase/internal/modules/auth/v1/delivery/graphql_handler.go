package delivery

import (
	"context"
	"fmt"

	"gitlab.com/Wuriyanto/go-codebase/internal/modules/auth/v1/usecase"
	"gitlab.com/Wuriyanto/go-codebase/pkg/middleware"
)

// GraphQLHandler model
type GraphQLHandler struct {
	authUsecase usecase.AuthUsecase
	mw          *middleware.Middleware
}

// NewGraphQLHandler delivery
func NewGraphQLHandler(uc usecase.AuthUsecase, mw *middleware.Middleware) *GraphQLHandler {
	return &GraphQLHandler{
		authUsecase: uc,
		mw:          mw,
	}
}

type (
	// LoginArgs args
	LoginArgs struct {
		Username, Password string
	}
	// UserSchema model
	UserSchema struct {
		ID                              int32
		Username, Password, Name, Token string
	}
)

// Login handler
func (h *GraphQLHandler) Login(ctx context.Context, args *LoginArgs) (*UserSchema, error) {
	fmt.Printf("%+v\n", args)
	return &UserSchema{}, nil
}
