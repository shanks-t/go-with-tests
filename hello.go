package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("Trey"))
}

func Hello(name string) string {
	if name == "" {
		name = "Worldd"
	}
	return englishHelloPrefix + name
}
