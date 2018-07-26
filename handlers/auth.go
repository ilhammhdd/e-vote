package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ilhammhdd/e-vote/models"
	"github.com/ilhammhdd/e-vote/utils"
)

func GenerateToken(w http.ResponseWriter, r *http.Request) {
	userClaims := models.User{
		Id:       99,
		Username: models.NullString{sql.NullString{r.URL.Query()["username"][0], true}},
		Email:    models.NullString{sql.NullString{r.URL.Query()["email"][0], true}},
	}

	token, _ := utils.GenerateSignedToken(userClaims)
	fmt.Fprintln(w, token)
}

func VerifyToken(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "token valid")
}
