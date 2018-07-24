package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ilhammhdd/e-vote/models"
	"github.com/ilhammhdd/e-vote/utils"
)

func GenerateToken(w http.ResponseWriter, r *http.Request) {
	userClaims := &UserClaims{
		models.User{
			Id:       99,
			Username: models.NullString{Value: sql.NullString{String: r.URL.Query()["username"][0], Valid: true}},
			Email:    models.NullString{Value: sql.NullString{String: r.URL.Query()["email"][0], Valid: true}},
		},
		jwt.StandardClaims{
			IssuedAt:  time.Now().UnixNano() / 1000000,
			ExpiresAt: (time.Now().UnixNano() / 1000000) + 21600000,
		},
	}

	token, err := utils.GenerateRSATokenWithClaims(*userClaims)
	handleError(err)
	fmt.Fprintln(w, token)
}

func VerifyToken(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "token valid")
}
