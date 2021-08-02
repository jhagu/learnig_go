package main

import "fmt"

const french = "FR"
const spanish = "ES"
const italian = "IT"
const englishPrefixHello = "Hello, "
const spanishPrefixHello = "Hola, "
const frenchPrefixHello = "Bonjour, "
const italianPrefixHello = "Ciao, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

//Private function with named return value (prefix)
// With a named return value, the return statement returns whatever is set, in this case the' prefix' variable

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefixHello
	case spanish:
		prefix = spanishPrefixHello
	case italian:
		prefix = italianPrefixHello
	default:
		prefix = englishPrefixHello
	}
	return
}

func main() {
	fmt.Println(Hello("Juan", "ES"))
	fmt.Println(Hello("", "IT"))
}
