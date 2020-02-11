package main

import (
	"fmt"
	"strings"
)

func main() {

	// Get the String Length.
	fmt.Println(len("GoLang")) // prints 6

	// Convert String to Title Case.
	fmt.Println(strings.Title("go is fun to learn.")) // Prints "Go Is Fun To Learn."

	// Check if String Starts with a SubString.
	fmt.Println(strings.HasPrefix("GoLang", "Go")) // prints true

	// Check if String Ends with SubString.
	fmt.Println(strings.HasSuffix("GoLang", "ng")) // prints true

	// Compare if Two Strings are Equal (Ignore Case).
	fmt.Println(strings.EqualFold("GOLANG", "golang")) // prints true

}
