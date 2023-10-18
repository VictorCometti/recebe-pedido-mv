package main

import (
	"application/pkg/router"
	"log"
	"net/http"
)

func main() {
	r := router.GetRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
