package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/context"
)

type middleware func(next http.HandlerFunc) http.HandlerFunc

type AuthResponse struct {
	Payload User `json:"payload"`
}

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logged connection from %s", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}

func withTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request for %s", r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

func checkAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("x-auth-token")
		log.Printf("Got Header: %s\n", token)
		if token == "" {
			log.Println("Token not provided, returning error")
			WriteError(w, http.StatusUnauthorized, errors.New("Not Authorized"))
			return
		}
		v := make(map[string]string)
		v["token"] = token
		log.Println("Validating Token")
		resp, err := Post("http://localhost:8080/validateToken", v)
		if err != nil {
			log.Println("Error validating Token")
			log.Println(err.Error())
			WriteError(w, http.StatusBadRequest, err)
			return
		}
		if resp.StatusCode != 200 {
			log.Println("Token Validated with error code")
			WriteError(w, http.StatusUnauthorized, errors.New("Not Authenticated"))
			return
		}
		var payload AuthResponse
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(&payload); err != nil {
			log.Println("Error decoding payload for authentication middleware")
			log.Fatal(err)
		}
		log.Println(payload)
		u := payload.Payload
		context.Set(r, "User", u)
		next.ServeHTTP(w, r)
	}
}

// chainMiddleware provides syntactic sugar to create a new middleware
// which will be the result of chaining the ones received as parameters.
func chainMiddleware(final http.HandlerFunc, mw ...middleware) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for i := len(mw) - 1; i >= 0; i-- {
			final = mw[i](final)
		}
		final(w, r)
	}
}
