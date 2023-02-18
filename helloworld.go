package main

import (
	"flag"
	"fmt"
)

type language string

func main() {
	var userLang string
	//using flag to capture user imput
	flag.StringVar(&userLang, "lang", "en", "The required language, e.g. en, heb...")
	// fills all flag occurances with the value provided
	flag.Parse()

	greeting := greet(language(userLang))
	fmt.Println(greeting)
}

var phrasebook = map[language]string{
	"el":  "Χαίρετε Κόσμε",     // Greek
	"en":  "Hello World",       // English
	"fr":  "Bonjor le monde",   // French
	"heb": "שלום עולם",         // Hebrew
	"ur":  "ہیلو دنیا",         // Urdu
	"vi":  "Xin chào Thế Giới", // Vietnamese
	"itl": "Choi",              // Italian
}

func greet(l language) string {
	// value // bool
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}

func ErrorMessage() {
	fmt.Println("Please choose one of the supported languages")
	for lang, _ := range phrasebook {
		fmt.Printf("%s\n", lang)
	}
}
