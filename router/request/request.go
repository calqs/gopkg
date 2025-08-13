package request

import "net/http"

type Request[DataT any] interface {
	GetData(*http.Request) (*DataT, error)
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
