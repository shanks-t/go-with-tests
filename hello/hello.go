package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const chineseHelloPrefix = "Meehow, "
const frenchHelloPrefix = "Bonjour, "

const spanish = "Spanish"
const chinese = "Chinese"
const french = "French"

func main() {
	fmt.Println(Hello("", "Chinese"))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	}

	return prefix + name
}
