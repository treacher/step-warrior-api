package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Error(ctx context.Context, w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(status)

	resp := &ErrorResponse{
		Code:    code,
		Message: message,
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		fmt.Println(err)
	}
}

var InternalServerError = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	Error(
		r.Context(),
		w,
		http.StatusInternalServerError,
		"internal_server_error",
		"Internal Server Error")
})
