package main

import (
	"fmt"
)

const (
	spanish string = "Spanish"
	french  string = "French"

	englishHelloPrefix string = "Hello, "
	spanishHelloPrefix string = "Hola, "
	frendchHelloPrefix string = "Bonjour, "
)

func Hello(name string, language string) string {
	var prefix string

	if name == "" {
		name = "world"
	}

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frendchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("", "English"))
}
