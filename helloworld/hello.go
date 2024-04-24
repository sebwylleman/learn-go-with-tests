package main

import (
	"fmt"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	const (
		englishHelloPrefix = "Hello, "
		spanishHelloPrefix = "Hola, "
		frenchHelloPrefix  = "Bonjour, "
	)
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
