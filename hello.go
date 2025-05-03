package main

import "fmt"

const englishPrefixString = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishPrefixString + name
}

func main() {
	fmt.Println(Hello("World"))
}
