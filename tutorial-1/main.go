package main

import "fmt"

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
		fmt.Printf("%x", str[i])
	}
}

func iterateStringSmart(str string) {
	for index, char := range str {
		fmt.Printf("%q %c", index, char)
	}
}

func main() {
	fmt.Println(Hello())
	fmt.Println(Hello2("Soma"))
	fmt.Println(maths(10.5, 10))
	iterateStringNaive("Hello")
	iterateStringSmart("Hello")
}
