package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type contextKey string

const validatedBodyKey cxtKey = "validatedBody"

var validate = validator.New()

func ValidateBody[T any]() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var body T

			// Decode JSON request
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				http.Error(w, "invalid JSON", http.StatusBadRequest)
				return
			}

			// Validate struct fields using tags
			if err := validate.Struct(body); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Add the validated body to context
			ctx := context.WithValue(r.Context(), validatedBodyKey, body)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetValidatedBody retrieves the validated DTO from context.
func GetValidatedBody[T any](r *http.Request) (T, bool) {
	val, ok := r.Context().Value(validatedBodyKey).(T)
	return val, ok
}
