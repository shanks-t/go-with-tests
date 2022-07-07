package main

import (
	dependencyinjection "di/test"
	"log"
	"net/http"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	dependencyinjection.Greet(w, "muhfucka")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
