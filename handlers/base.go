package handlers

import (
	"log"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/ilhammhdd/e-vote/models"
)

var (
	f models.File
	c models.Category
)

type UserClaims struct {
	models.User `json:"user_data"`
	jwt.StandardClaims
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
}
