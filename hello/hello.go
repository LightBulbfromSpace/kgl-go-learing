package main

import "fmt"

const englishHelloPrefix = "Hello, "
const russianHelloPrefix = "Привет, "
const germanHelloPrefix = "Hallo, "
const englishWorld = "World"
const russianWorld = "Мир"
const germanWorld = "Welt"
const russian = "Russian"
const german = "German"

func Hello(name string, language string) string {

	if name == "" {
		name = DefaultReceiver(language)
	}
	return GreetingPrefix(language) + name
}

func GreetingPrefix(language string) (prefix string) {
	switch language {
	case russian:
		prefix = russianHelloPrefix
	case german:
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func DefaultReceiver(language string) (receiver string) {
	switch language {
	case russian:
		receiver = russianWorld
	case german:
		receiver = germanWorld
	default:
		receiver = englishWorld
	}
	return
}

func main() {
	fmt.Println(Hello("User", "Russian"))
}
