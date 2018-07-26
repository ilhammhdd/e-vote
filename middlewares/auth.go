package middlewares

import (
	"fmt"
	"net/http"

	"github.com/ilhammhdd/e-vote/utils"
)

func Authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !utils.VerifySignedToken(r.Header.Get("Token")) {
			fmt.Fprintln(w, "Invalid token")
			return
		}
		h.ServeHTTP(w, r)
	})
}
