// Package token provides an http middleware for http basic auth token
// validation.
package token

import "net/http"

// ValidateTokens will return a handler that will verify that a session
// exists before allowing the handler in the arugment to be called.
// If the TokenValidator returns false it responds with a 401 code.
func ValidateTokens(v TokenValidator, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if token := getToken(r); !v.ValidToken(token) {
			Unauthorized(w)
			return
		}

		h.ServeHTTP(w, r)
	})
}

// Unauthorized is just a convience function that allows us to write a
// status code of 401 and a message of "Unauthorized" to the response
func Unauthorized(w http.ResponseWriter) {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

// getToken pulls the token from the Basic Auth Header.
// It assumes that it is the username
func getToken(r *http.Request) (token string) {
	token, _, _ = r.BasicAuth()
	return token
}
