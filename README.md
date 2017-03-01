# Token

Token is a simple middleware that reads a token from the basic auth header. It assumes that the token is stored in the user name field.

## Example
```go
	var (
		validator = ValidatorFunc(func(token string) bool {
			return token == "SECRET"
		})
		handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	)

	// Verifies token == "SECRET" before calling to handler
	verifiedHandler := VerifyTokens(validator, handler)

	log.Fatal(http.ListenAndServe(":8080", verifiedHandler))
```


