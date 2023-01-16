package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const korean = "Korean"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const koreanHelloPrefix = "Annyeonghaseyo, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case korean:
		prefix = koreanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {

	fmt.Println(Hello("world", ""))
}
