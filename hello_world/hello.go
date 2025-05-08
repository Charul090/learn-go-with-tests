package main

import (
	"fmt"
)

const (
	englishPrefixString = "Hello, "
	spanishPrefixString = "Hola, "

	frenchPrefixString = "Bonjour, "
	spanishString      = "Spanish"
	frenchString       = "French"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return generatePrefix(language) + name
}

func generatePrefix(language string) (prefix string) {
	switch language {
	case spanishString:
		prefix = spanishPrefixString
	case frenchString:
		prefix = frenchPrefixString
	default:
		prefix = englishPrefixString
	}

	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
