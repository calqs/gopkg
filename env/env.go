package env

import (
	"os"

	"github.com/calqs/gopkg/dt"
)

var EnvTagName = "env"

func ParseEnv[EnvStructT any]() (EnvStructT, error) {
	return dt.DynamicParseStruct[EnvStructT](EnvTagName, func(s string) string { return os.Getenv(s) })
}
