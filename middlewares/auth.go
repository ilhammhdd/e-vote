package middlewares

import (
	"fmt"
	"net/http"

	"github.com/ilhammhdd/e-vote/utils"
)

func Authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if err := utils.VerifyRSAFromPEM(r.Header.Get("Token")); err != nil {
			fmt.Fprintln(w, "invalid token")
			fmt.Fprintln(w, err)
			return
		}
		h.ServeHTTP(w, r)
	})
}
