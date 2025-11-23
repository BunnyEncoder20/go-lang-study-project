package main

import (
	"fmt"
	"strings"
)

func appendingToString() {
	strSlice := []string{"s", "o", "m", "y", "a"}
	catStr := ""

	// Ieffecient way to concatenating strings (if we are doing a lot of it)
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr)

	// Performant way: Using the built-in "strings" package's stringBuilder()
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	fmt.Printf("\n%v", strBuilder.String())
}

func main() {
	myString := "résumé"
	fmt.Println(len(myString))
	for idx, char := range myString {
		fmt.Println(idx, char)
	}

	myRune := []rune("résumé")
	fmt.Println(len(myString))
	for idx, char := range myRune {
		fmt.Println(idx, char)
	}

	appendingToString()
}
