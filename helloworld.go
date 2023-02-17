package main

import "fmt"

type language string

func main() {
   greeting := greet("heb")
   fmt.Println(greeting)
}

func greet(l language) string {
    switch l  {
    case "en":
        return "Hello World"
    case "fr":
        return "Bonjour le monde"
    case "heb":
        return "Shalom Olam"
    default:
        return ""
    }
}
