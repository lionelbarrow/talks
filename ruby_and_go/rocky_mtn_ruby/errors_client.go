package main

import (
	"fmt"
	"strconv"
)

func printParseResult(input string) {
	result, err := strconv.ParseBool(input)
	if err != nil {
		fmt.Printf("Unable to parse %s: %s\n", input, err.Error())
	} else {
		fmt.Printf("Parsed '%s' successfully: %v\n", input, result)
	}
}

func main() {
	// Parsed 'true' successfully: true
	printParseResult("true")

	// Unable to parse tru: strconv.ParseBool: parsing "tru": invalid syntax
	printParseResult("tru")
}
