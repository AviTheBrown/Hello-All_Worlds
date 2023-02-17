package main

import (
    "testing"
)

//func ExampleMain(){
//   main()
//   // Output:
//   // Hello World
//}

func TestGreet_English(t *testing.T) {
    lang := "en"
    expectedGreeting := "Hello World"
    greeting := greet(language(lang))
    if greeting != expectedGreeting {
        // mark this a failed
        t.Errorf("expected %q, got: %q", expectedGreeting, greeting)
    }
}
func TestGreet_French(t *testing.T) {
    lang := "fr"
    expectedGreeting := "Bonjour le monde"
    greeting := greet(language(lang))

    if expectedGreeting != greeting {
        t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
    }
}

func TestGreeting_Hebrew(t *testing.T)  {
    lang := "heb"
    expectedGreeting := "Shalom Olam"
    greeting := greet(language(lang))

    if expectedGreeting != greeting {
        t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
    }
}