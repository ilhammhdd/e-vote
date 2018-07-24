package handlers

import (
	"log"
	"net/http"
)

func IndexHome(w http.ResponseWriter, r *http.Request) {
	log.Println("testing duls")
	log.Println(r)
}
