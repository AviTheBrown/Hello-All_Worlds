package main

import "fmt"

type language string

func main() {
	greeting := greet("heb")
	fmt.Println(greeting)
}

var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello World",       // English
	"fr": "Bonjor le monde",   // French
	"he": "שלום עולם",         // Hebrew
	"ur": "ہیلو دنیا",         // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
}

func greet(l language) string {
	// value // bool
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported langiage: %q", l)
	}
	return greeting
}
