package dt

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
		trial, err := DynamicParseStruct[testStruct]("env", func(tag string) string { return os.Getenv(tag) })
		assert.NoError(t, err)
		assert.Equal(t, testStruct{123}, trial)
	})
	t.Run("with a default value", func(t *testing.T) {
		type testStruct struct {
			Cabane int `env:"cabane,?123"`
		}
		trial, err := DynamicParseStruct[testStruct]("env", func(tag string) string { return os.Getenv(tag) })
		assert.NoError(t, err)
		assert.Equal(t, testStruct{123}, trial)
	})
	t.Run("with uint", func(t *testing.T) {
		os.Setenv("cabane", "123")
		os.Setenv("cabane8", "123")
		os.Setenv("cabane16", "12345")
		os.Setenv("cabane32", "123456")
		os.Setenv("cabane64", "13347235300527625959")
		type testStruct struct {
			Cabane   uint   `env:"cabane"`
			Cabane8  uint8  `env:"cabane8"`
			Cabane16 uint16 `env:"cabane16"`
			Cabane32 uint32 `env:"cabane32"`
			Cabane64 uint64 `env:"cabane64"`
		}
		trial, err := DynamicParseStruct[testStruct]("env", func(tag string) string { return os.Getenv(tag) })
		assert.NoError(t, err)
		goal := testStruct{
			Cabane:   123,
			Cabane8:  123,
			Cabane16: 12345,
			Cabane32: 123456,
			Cabane64: 13347235300527625959,
		}
		assert.Equal(t, goal, trial)
	})
}
