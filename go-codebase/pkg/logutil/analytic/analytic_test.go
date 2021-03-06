package analytic

import (
	"os"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicAuth(t *testing.T) {

	t.Run("Test init analytic", func(t *testing.T) {
		InitAnalytic(os.Stdout)

		assert.NotNil(t, Anaytic)
		Anaytic = nil
	})

	t.Run("Test Log analytic", func(t *testing.T) {
		InitAnalytic(os.Stdout)

		data := &Data{}

		data.Message = "Log test"

		Log(data)

		assert.Equal(t, "user-service", data.Label)

		assert.NotNil(t, Anaytic)

		Anaytic = nil
	})

	t.Run("Negative Test Log analytic", func(t *testing.T) {

		data := &Data{}

		data.Message = "Log test"

		assert.Panics(t, func() {
			Log(data)
		})
	})

}
