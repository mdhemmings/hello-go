package main

import (
	"errors"
	"fmt"
	"os"
)

const prefix = "Hello "
const helloNumber = 5
const spanishHello = "Hola "
const spanish = "Spanish"
const french = "French"
const frenchHello = "Bonjour "
const englishHello = "Hello "

// PrefixSet : sets the prefix into locale specified

func prefixSet(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHello
	case spanish:
		prefix = spanishHello
	default:
		prefix = englishHello
	}
	return

}

// Hello : sets the hello world string
func hello(name string, language string) (string, error) {
	/*	if name == "" {
				name = "world"
		}
	*/
	if name == "" {
		return "", errors.New("argh")
	}

	return prefixSet(language) + name, nil
}

// Number : sets the number string
func number() string {
	return "5"
}

// Main : Combines the hello and number functions for fun and profit
func main() {

	helloMessage, err := hello("Bob", prefix)
	if err != nil {
		os.Exit(-1)
	}

	fmt.Println(helloMessage, helloNumber)
}
