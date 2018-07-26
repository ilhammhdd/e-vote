package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ilhammhdd/e-vote/handlers"
	"github.com/ilhammhdd/e-vote/middlewares"
	"github.com/ilhammhdd/e-vote/utils"
)

func main() {
	utils.OpenDB("root", "", "", "e_vote")
	defer utils.DB.Close()

	utils.GenerateECDSAKeyPairToFile()

	http.Handle("/generate/token", middlewares.MustParams(http.HandlerFunc(handlers.GenerateToken), "username", "email"))
	http.Handle("/verify/token", middlewares.MustHeaderParams(middlewares.Authenticate(http.HandlerFunc(handlers.VerifyToken)), "Token"))

	srv := &http.Server{
		Addr:         ":9090",
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 7 * time.Second,
	}

	log.Println(srv.ListenAndServe())
}
