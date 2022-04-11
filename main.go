package main

import (
	"log"
	"net/http"

	"github.com/TamirlanKadyrbekov/midterm/routes"
)

func main() {

	router := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
