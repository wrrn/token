# Token

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](http://godoc.org/github.com/wrrn/token)
[![Build Status](https://travis-ci.org/wrrn/token.svg?branch=master)](https://travis-ci.org/wrrn/token)
[![Coverage](http://gocover.io/_badge/github.com/wrrn/token)](http://gocover.io/github.com/wrrn/token)


Token is a simple middleware that reads a token from the basic auth header. It assumes that the token is stored in the user name field.

## Usage

The key function to the tokens package is ```ValidateTokens()```. ```ValidateTokens()``` takes anything that implements ```TokenValidator```. 
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
	validatedHandler := ValidateTokens(validator, handler)

	log.Fatal(http.ListenAndServe(":8080", validatedHandler))
```

### ValidatorFunc
 Use a function as the TokenValidator by casting a function to the appropriate signature.

```go
 ValidateTokens(func(token string) bool {
     return token == "SECRET"
 }, handler)
 ```


