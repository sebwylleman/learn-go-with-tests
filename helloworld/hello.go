package main

import "fmt"

const language = "Spanish"
const englishHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == language {
		return SpanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
