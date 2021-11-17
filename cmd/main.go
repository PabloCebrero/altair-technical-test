package main

import (
	"fmt"
	"github.com/altair-tecnical-test/src/server"
	"log"
	"net/http"
)
func main() {
	s := server.New()

	fmt.Printf("Listening port 8080...")
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
