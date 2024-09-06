package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Encode[T any](w http.ResponseWriter, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func DecodeJSONBody[T any](r *http.Request, v *T) error {
	return DecodeJSON[T](r, v)
}

func DecodeJSON[T any](r *http.Request, v *T) error {
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return fmt.Errorf("decode json: %w", err)
	}
	return nil
}
