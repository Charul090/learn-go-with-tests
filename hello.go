package main

import "fmt"

const englishPrefixString = "Hello, "
const spanishPrefixString = "Hola, "
const spanishString = "Spanish"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanishString {
		return spanishPrefixString + name
	}
	return englishPrefixString + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
