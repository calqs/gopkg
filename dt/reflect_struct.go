package dt

import (
	"reflect"
	"strconv"
	"strings"
)

var DefaultKeyChar byte = '?'

type Tag struct {
	actualValue  *string
	defaultValue *string
}

func NewTag(rawTag string) Tag {
	parts := strings.Split(rawTag, ",")
	tag := Tag{}
	if len(parts) == 0 {
		return tag
	}
	tag.actualValue = PtrNilOnEmpty(parts[0])
	if len(parts) > 1 {
		parts = parts[1:]
	}
	for _, part := range parts {
		if len(part) < 1 {
			continue
		}
		switch part[0] {
		case DefaultKeyChar:
			tag.defaultValue = PtrNilOnEmpty(part[1:])
		}
	}
	return tag
}

func assignValue(i int, val string, v reflect.Value) {
	fv := v.Field(i)
	switch fv.Kind() {
	case reflect.String:
		fv.SetString(val)
	case reflect.Int:
		if n, err := strconv.Atoi(val); err == nil {
			fv.SetInt(int64(n))
		}
	case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint:
		if n, err := strconv.ParseUint(val, 0, 64); err == nil {
			fv.SetUint(uint64(n))
		}
		// @TODO: add more types
	}
}

func DynamicParseStruct[T any](tagName string, matcher func(string) string) (T, error) {
	var dst T
	v := reflect.ValueOf(&dst).Elem()
	t := v.Type()

	for i := range t.NumField() {
		field := t.Field(i)
		rawTag := field.Tag.Get(tagName)
		if rawTag == "" {
			rawTag = strings.ToLower(field.Name)
		}
		tag := NewTag(rawTag)

		if val := matcher(Deref(tag.actualValue)); tag.actualValue != nil && val != "" {
			assignValue(i, val, v)
			continue
		}
		if tag.defaultValue != nil {
			assignValue(i, Deref(tag.defaultValue), v)
		}
	}
	return dst, nil
}
