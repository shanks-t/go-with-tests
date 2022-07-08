package main

import (
	dependencyinjection "di/test"
	mocking "mocking/countdown"
	"net/http"
	"os"
)

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	dependencyinjection.Greet(w, "muhfucka")
}

func main() {
	//log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))

	mocking.Countdown(os.Stdout)
}
