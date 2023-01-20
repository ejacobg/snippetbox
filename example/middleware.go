package example

import (
	"net/http"
)

func middleware1(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Middleware logic here...
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Middleware logic here...
		next.ServeHTTP(w, r)
	})
}
