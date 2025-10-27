package req

import (
	"link-shorter/pkg/res"
	"net/http"
)

func HandleBody[T any](w http.ResponseWriter, req *http.Request) (*T, error) {
	payload, err := Decoder[T](req.Body)

	if err != nil {
		res.Json(w, err.Error(), 402)
		return nil, err
	}

	err = isValid(payload)
	if err != nil {
		res.Json(w, err.Error(), 402)
		return nil, err
	}

	return &payload, nil
}