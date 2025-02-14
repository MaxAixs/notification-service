package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func parseJSONBody(w http.ResponseWriter, r *http.Request, input interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(input); err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
		case errors.As(err, &unmarshalTypeError):
			http.Error(w, "Invalid value", http.StatusBadRequest)
		case errors.Is(err, io.EOF):
			http.Error(w, "Empty body", http.StatusBadRequest)
		default:
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
		}

		return err
	}

	return nil
}
