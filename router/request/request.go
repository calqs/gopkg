package request

import (
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func setField(field reflect.Value, value string) error {
	// If it's a pointer, allocate and move into the element
	if field.Kind() == reflect.Pointer {
		// Create a new instance of the underlying type
		// e.g., if field is *int, result is a reflect.Value of an int pointer
		ptr := reflect.New(field.Type().Elem())

		// Recursively call setField on the element the pointer points to
		if err := setField(ptr.Elem(), value); err != nil {
			return err
		}

		// Set the struct field to our new pointer
		field.Set(ptr)
		return nil
	}

	// Actual data conversion logic

	// Specific type handling
	switch field.Type() {
	case reflect.TypeFor[time.Time]():
		t, err := time.Parse(time.RFC3339, value)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(t))
		return nil
	}

	// Primitive type handling
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(i)
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		field.SetBool(b)
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		field.SetFloat(f)
	}
	return nil
}

// bindValues tries to find "query" tag in the passed pointer to generic type
// and tries to assign a querystring parameter matching that tag
func bindValues[T any](dst *T, vals url.Values) error {
	v := reflect.ValueOf(dst).Elem()
	t := v.Type()

	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get("query")
			if tag == "" {
				tag = strings.ToLower(field.Name)
			}

			tagVals, exists := vals[tag]
			if exists && len(tagVals) > 0 {
				// tagVals[0] could be "", but we still want to set it
				// especially if the struct field is a *string
				if err := setField(v.Field(i), tagVals[0]); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// ExtractData will retrieve and transform data from the request's querystring and body.
// @todo: make querystring retrieval explici
func ExtractData[DataT any](r *http.Request) (*DataT, error) {
	res, err := JsonBodyRequest[DataT](r)
	if err != nil {
		return nil, err
	}
	if res == nil {
		var d DataT
		res = &d
	}
	err = bindValues(res, r.URL.Query())
	if err != nil {
		return nil, err
	}
	return res, nil
}
