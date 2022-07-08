package main

import (
	mocking "mocking/countdown"
	"os"
)

// func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
// 	dependencyinjection.Greet(w, "muhfucka")
// }

func main() {
	//log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
	sleeper := &mocking.DefaultSleeper{}
	mocking.Countdown(os.Stdout, sleeper)
}
