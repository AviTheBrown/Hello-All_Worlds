package main

import (
	"testing"
)

//func ExampleMain(){
//   main()
//   // Output:
//   // Hello World
//}

func TestGreet(t *testing.T) {
	// prep phase
	type scenario struct {
		lang             language
		expectedGreeting string
	}
	var tests = map[string]scenario{

		"English": {
			lang:             "en",
			expectedGreeting: "Hello World",
		},
		"French": {
			lang:             "fr",
			expectedGreeting: "Bonjor le monde",
		},
		"Akkadian, not supported": {
			lang:             "akk",
			expectedGreeting: "unsupported language: \"akk\"",
		},
		"Greek": {
			lang:             "el",
			expectedGreeting: "Χαίρετε Κόσμε",
		},
		"Hebrew": {
			lang:             "heb",
			expectedGreeting: "שלום עולם",
		},
		"Urdu": {
			lang:             "ur",
			expectedGreeting: "ہیلو دنیا",
		},
		"Vietnamese": {
			lang:             "vi",
			expectedGreeting: "Xin chào Thế Giới",
		},
		"Empty": {
			lang:             "",
			expectedGreeting: "unsupported language: \"\"",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// execution phase call function
			greeting := greet(tc.lang)
			// decision phase
			if greeting != tc.expectedGreeting {
				t.Errorf("expected: %q, got: %q", tc.expectedGreeting, greeting)
			}
		})
	}
}
