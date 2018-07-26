package handlers

import (
	"log"

	"github.com/ilhammhdd/e-vote/models"
)

var (
	f models.File
	c models.Category
)

func handleError(err error) {
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
}
