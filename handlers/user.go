package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func UserIndex(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	fmt.Fprint(w, "Written response")
}
