package request

import (
	"errors"
	"fmt"
)

var (
	ErrNilPointer            = errors.New("nil pointer")
	ErrPayloadMalformed      = errors.New("payload malformed")
	ErrPayloadWrongFieldType = fmt.Errorf("%w: wrong field type", ErrPayloadMalformed)
	ErrPayloadWrongShape     = fmt.Errorf("%w: does not match a json payload shape", ErrPayloadMalformed)
)
