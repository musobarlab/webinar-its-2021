package token

import (
	"context"
	"crypto/rsa"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"gitlab.com/Wuriyanto/go-codebase/pkg/helper"
	"gitlab.com/Wuriyanto/go-codebase/pkg/shared"
)

// JWT repo
type JWT struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// NewJWT constructor
func NewJWT(publicKey *rsa.PublicKey, privateKey *rsa.PrivateKey) *JWT {
	return &JWT{
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

// Generate token
func (r *JWT) Generate(ctx context.Context, payload *Claim, expired time.Duration) <-chan shared.Result {
	output := make(chan shared.Result)

	go func() {
		defer close(output)

		now := time.Now()
		exp := now.Add(expired)

		var key interface{}
		var token = new(jwtgo.Token)
		if payload.Alg == HS256 {
			token = jwtgo.New(jwtgo.SigningMethodHS256)
			key = []byte(helper.TokenKey)
		} else {
			token = jwtgo.New(jwtgo.SigningMethodRS256)
			key = r.privateKey
		}
		claims := jwtgo.MapClaims{
			"iss": "telkomdev",
			"exp": exp.Unix(),
			"iat": now.Unix(),
			"sub": payload.User.ID,
			"aud": "97b33193-43ff-4e58-9124-b3a9b9f72c34",
			"jti": payload.JTI,
		}
		if payload.User.Email != "" {
			claims["email"] = payload.User.Email
		}
		if payload.RoleUser != "" {
			claims["roleUser"] = payload.RoleUser
		}
		if payload.User.AppID != "" {
			claims["appId"] = payload.User.AppID
		}
		if payload.DeviceID != "" {
			claims["did"] = payload.DeviceID
		}
		if payload.RefreshJTI != "" {
			claims["refresh"] = payload.RefreshJTI
		}

		token.Claims = claims

		tokenString, err := token.SignedString(key)
		if err != nil {
			output <- shared.Result{Error: err}
			return
		}

		output <- shared.Result{Data: tokenString}
	}()

	return output
}

// Refresh token
func (r *JWT) Refresh(ctx context.Context, token string) <-chan shared.Result {
	output := make(chan shared.Result)

	go func() {
		defer close(output)
	}()

	return output
}

// Validate token
func (r *JWT) Validate(ctx context.Context, tokenString string) <-chan shared.Result {
	output := make(chan shared.Result)

	go func() {
		defer close(output)

		tokenParse, err := jwtgo.Parse(tokenString, func(token *jwtgo.Token) (interface{}, error) {
			checkAlg, _ := shared.GetValueFromContext(ctx, shared.ContextKey("tokenAlg")).(string)
			if checkAlg == HS256 {
				return []byte(helper.TokenKey), nil
			}
			return r.publicKey, nil
		})

		var errToken error
		switch ve := err.(type) {
		case *jwtgo.ValidationError:
			if ve.Errors == jwtgo.ValidationErrorExpired {
				errToken = helper.ErrTokenExpired
			} else {
				errToken = helper.ErrTokenFormat
			}
		}

		if errToken != nil {
			output <- shared.Result{Error: errToken}
			return
		}

		if !tokenParse.Valid {
			output <- shared.Result{Error: helper.ErrTokenFormat}
			return
		}

		mapClaims, _ := tokenParse.Claims.(jwtgo.MapClaims)

		var tokenClaim Claim
		tokenClaim.DeviceID, _ = mapClaims["did"].(string)
		tokenClaim.RoleUser, _ = mapClaims["roleUser"].(string)
		tokenClaim.User.ID, _ = mapClaims["sub"].(string)
		tokenClaim.User.Email, _ = mapClaims["email"].(string)
		tokenClaim.JTI, _ = mapClaims["jti"].(string)
		tokenClaim.RefreshJTI, _ = mapClaims["refresh"].(string)

		output <- shared.Result{Data: &tokenClaim}
	}()

	return output
}

// Revoke token
func (r *JWT) Revoke(ctx context.Context, token string) <-chan shared.Result {
	output := make(chan shared.Result)

	go func() {
		defer close(output)
	}()

	return output
}
