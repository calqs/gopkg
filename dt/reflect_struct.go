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

func assignValue(v reflect.Value, val string) {
	if v.Kind() == reflect.Pointer {
		ptr := reflect.New(v.Type().Elem())
		assignValue(ptr.Elem(), val)
		v.Set(ptr)
		return
	}

	switch v.Kind() {
	case reflect.String:
		v.SetString(val)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if n, err := strconv.ParseInt(val, 10, 64); err == nil {
			v.SetInt(n)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if n, err := strconv.ParseUint(val, 10, 64); err == nil {
			v.SetUint(n)
		}
	case reflect.Bool:
		if b, err := strconv.ParseBool(val); err == nil {
			v.SetBool(b)
		}
		// @TODO: Add Float, etc.
	}
}

func DynamicParseStruct[T any](tagName string, matcher func(string) string) (T, error) {
	var dst T
	v := reflect.ValueOf(&dst).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fv := v.Field(i)
		rawTag := field.Tag.Get(tagName)
		if rawTag == "" {
			rawTag = strings.ToLower(field.Name)
		}
		tag := NewTag(rawTag)

		matchKey := Deref(tag.actualValue)
		val := matcher(matchKey)

		if val != "" {
			assignValue(fv, val)
		} else if tag.defaultValue != nil {
			assignValue(fv, Deref(tag.defaultValue))
		}
	}
	return dst, nil
}
