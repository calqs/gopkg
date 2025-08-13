package request

import (
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func bindValues[T any](dst *T, vals url.Values) error {
	v := reflect.ValueOf(dst).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("query")
		if tag == "" {
			tag = strings.ToLower(field.Name)
		}

		if val := vals.Get(tag); val != "" {
			fv := v.Field(i)
			switch fv.Kind() {
			case reflect.String:
				fv.SetString(val)
			case reflect.Int:
				if n, err := strconv.Atoi(val); err == nil {
					fv.SetInt(int64(n))
				}
				// @TODO: add more types
			}
		}
	}
	return nil
}

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
