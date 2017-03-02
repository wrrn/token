package token

import (
	"log"
	"net/http"
)

func ExampleValidateTokens() {
	var (
		validator = ValidatorFunc(func(token string) bool {
			return token == "SECRET"
		})
		handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	)

	// Verifies token == "SECRET" before calling to handler
	validatedHandler := ValidateTokens(validator, handler)

	log.Fatal(http.ListenAndServe(":8080", validatedHandler))
}
