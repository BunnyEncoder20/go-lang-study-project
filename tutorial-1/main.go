package main

import (
	"errors"
	"fmt"
	"log"
)

func Hello() string {
	return "Hello World"
}

const englishGreetingPrefix = "Hello, "

func Hello2(name string) string {
	if name == "" {
		return Hello()
	}

	return englishGreetingPrefix + name
}

func maths(num1 float32, num2 int32) float32 {
	results := num1 + float32(num2)
	return results
}

func iterateStringNaive(str string) {
	// Prints bytes: 72, 105, 206, 147 (Garbage at the end)
	for i := 0; i < len(str); i++ {
		fmt.Printf("%v\n", str[i])
	}
}

func iterateStringSmart(str string) {
	for index, char := range str {
		fmt.Printf("%v %c\n", index, char)
	}
}

func myDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	fmt.Println(Hello())
	fmt.Println(Hello2("Soma"))
	fmt.Println(maths(10.5, 10))
	iterateStringNaive("Hello")
	iterateStringSmart("Hello")

	results, remainder, err := myDivision(10, 2)
	if err != nil {
		log.Println(err)
	} else if remainder == 0 {
		fmt.Printf("results: %v\n", results)
	} else {
		fmt.Printf("results: %v, remainder: %v \n", results, remainder)
	}

	results, remainder, err = myDivision(10, 3)
	switch {
	case err != nil:
		log.Println(err)
	case remainder == 0:
		fmt.Printf("results: %v\n", results)
	default:
		fmt.Printf("results: %v, remainder: %v \n", results, remainder)
	}

	// Arrays
	var intArr [3]int32 // declaration and later assignment
	intArr[0] = 1
	intArr[1] = 2
	intArr[2] = 3

	fmt.Println(intArr[0])
	fmt.Println(intArr[0:2])
	fmt.Println(intArr[1:3])

	intArr2 := [3]int16{1, 2, 3} // declaration and assignment
	fmt.Println(intArr2)
}
