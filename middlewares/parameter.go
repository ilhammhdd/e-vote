package middlewares

import (
	"fmt"
	"net/http"
)

var missingParams []string

func MustParams(h http.Handler, params ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		for _, param := range params {
			if _, ok := q[param]; !ok {
				missingParams = append(missingParams, param)
			}
		}
		if len(missingParams) != 0 {
			for _, param := range missingParams {
				fmt.Fprintln(w, "missing param : ", param)
			}
			missingParams = missingParams[:0]
			return
		}
		h.ServeHTTP(w, r)
	})
}

func MustHeaderParams(h http.Handler, params ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, param := range params {
			if token := r.Header.Get(param); token == "" {
				missingParams = append(missingParams, param)
			}
		}
		if len(missingParams) != 0 {
			for _, param := range missingParams {
				fmt.Fprintln(w, "missing header param : ", param)
			}
			missingParams = missingParams[:0]
			return
		}
		h.ServeHTTP(w, r)
	})
}
