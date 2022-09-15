package website

import (
	"log"
	"net/http"
)

func Run() {

	http.HandleFunc("/", indexHandle())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
