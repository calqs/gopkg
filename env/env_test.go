package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_I_Can_Parse_Struct_And_Apply_Values(t *testing.T) {
	t.Run("with a provided env var", func(t *testing.T) {
		type testStruct struct {
			Cabane int `env:"cabane,?123"`
		}
		os.Setenv("cabane", "123")
		t.Cleanup(func() {
			os.Unsetenv("cabane")
		})
		trial, err := ParseEnv[testStruct]()
		assert.NoError(t, err)
		assert.Equal(t, testStruct{123}, trial)
	})
	t.Run("with a default value", func(t *testing.T) {
		type testStruct struct {
			Cabane int `env:"cabane,?123"`
		}
		trial, err := ParseEnv[testStruct]()
		assert.NoError(t, err)
		assert.Equal(t, testStruct{123}, trial)
	})
}
