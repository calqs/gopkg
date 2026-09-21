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

func Test_I_Can_Parse_Struct_With_Byte_Arrays(t *testing.T) {
	t.Run("with a valid hex-encoded env var", func(t *testing.T) {
		os.Setenv("key32", "0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20")
		t.Cleanup(func() {
			os.Unsetenv("key32")
		})
		type testStruct struct {
			Key [32]byte `env:"key32"`
		}
		trial, err := DynamicParseStruct[testStruct]("env", func(tag string) string { return os.Getenv(tag) })
		assert.NoError(t, err)
		goal := testStruct{}
		for i := range goal.Key {
			goal.Key[i] = byte(i + 1)
		}
		assert.Equal(t, goal, trial)
	})
	t.Run("with a shorter env var than what is expected (leaves the array zeroed)", func(t *testing.T) {
		os.Setenv("key32", "01020304")
		t.Cleanup(func() {
			os.Unsetenv("key32")
		})
		type testStruct struct {
			Key [32]byte `env:"key32"`
		}
		trial, err := DynamicParseStruct[testStruct]("env", func(tag string) string { return os.Getenv(tag) })
		assert.NoError(t, err)
		assert.Equal(t, testStruct{}, trial)
	})
	t.Run("with a malformed env var, leaves the array zeroed", func(t *testing.T) {
		os.Setenv("key32", "not-hex")
		t.Cleanup(func() {
			os.Unsetenv("key32")
		})
		type testStruct struct {
			Key [32]byte `env:"key32"`
		}
		trial, err := DynamicParseStruct[testStruct]("env", func(tag string) string { return os.Getenv(tag) })
		assert.NoError(t, err)
		assert.Equal(t, testStruct{}, trial)
	})
}

func Test_I_Can_Parse_Struct_With_Pointers(t *testing.T) {
	t.Run("with a provided env var", func(t *testing.T) {
		type testStruct struct {
			Cabane *int `env:"cabane,?123"`
		}
		os.Setenv("cabane", "123")
		t.Cleanup(func() {
			os.Unsetenv("cabane")
		})
		trial, err := DynamicParseStruct[testStruct]("env", func(tag string) string { return os.Getenv(tag) })
		assert.NoError(t, err)
		assert.Equal(t, testStruct{Ptr(123)}, trial)
	})
	t.Run("with a default value", func(t *testing.T) {
		type testStruct struct {
			Cabane *int `env:"cabane,?123"`
		}
		trial, err := DynamicParseStruct[testStruct]("env", func(tag string) string { return os.Getenv(tag) })
		assert.NoError(t, err)
		assert.Equal(t, testStruct{Ptr(123)}, trial)
	})
}
