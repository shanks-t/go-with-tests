package main

import (
	mocking "mocking/countdown"
	"os"
	"time"
)

// func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
// 	dependencyinjection.Greet(w, "muhfucka")
// }

func main() {
	//log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
	sleeper := &mocking.ConfigurableSleeper{Duration: 1 * time.Second, ConfiguredSleep: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
