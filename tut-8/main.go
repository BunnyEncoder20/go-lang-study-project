package main

import "fmt"

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(sumSlice(intSlice))
}

func sumSlice[T int | float32 | float64](slice []T) {
	var sum T
	for v := range slice {
		sum += slice[v]
	}
	return sum
}
