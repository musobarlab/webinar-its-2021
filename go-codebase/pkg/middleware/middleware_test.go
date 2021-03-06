package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/Wuriyanto/go-codebase/config"
)

func TestNewMiddleware(t *testing.T) {
	mw := NewMiddleware(&config.Config{})
	assert.NotNil(t, mw)
}
