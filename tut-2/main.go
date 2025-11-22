package main

import (
	"fmt"
)

func showcaseArrays() {
	var intArr [3]int32 // declaration and later assignment
	intArr[0] = 1
	intArr[1] = 2
	intArr[2] = 3

	fmt.Println(intArr[0])
	fmt.Println(intArr[0:2])
	fmt.Println(intArr[1:3])

	intArr2 := [3]int16{1, 2, 3} // declaration and assignment
	fmt.Println(intArr2)

	for i, v := range intArr {
		fmt.Printf("intArr - Index: %v, value: %v \n", i, v)
	}
}

func showcaseSlices() {
	intSlice := []int32{4, 5, 6} // omit the length to make a slice
	fmt.Printf("intSlice Legnth: %v | Capacity: %v \n", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7) // Can append elements to slice, but if len == cap then a new array (with double the cap) is made and all previous elements are copied over
	fmt.Printf("intSlice Legnth: %v | Capacity: %v \n", len(intSlice), cap(intSlice))

	intSlice1 := []int32{10, 11, 12}
	intSlice2 := []int32{13, 14, 15}
	intSlice1 = append(intSlice1, intSlice2...) // can add multiple elements using the spread op
	fmt.Printf("intSlice1 Legnth: %v | Capacity: %v \n", len(intSlice1), cap(intSlice1))
	for i, v := range intSlice1 {
		fmt.Printf("intSlice1 - Index: %v, value: %v \n", i, v)
	}

	// For better performance and avoiding frequent reallocation, use the make() func
	intMakeSlice := make([]int32, 0, 10) // int32 slice with len 0 and cap 10
	fmt.Printf("intMakeSlice Legnth: %v | Capacity: %v \n", len(intMakeSlice), cap(intMakeSlice))
}

func showcaseMaps() {
	myMap1 := make(map[string]int)
	myMap1["Varun"] = 8
	myMap1["Soma"] = 8
	fmt.Println(myMap1)

	myMap2 := map[string]int{"Varun": 8, "Soma": 8, "Tarun": 10}
	fmt.Println(myMap2["Soma"]) // output: 8
	fmt.Println(myMap2["Chaitanya"])

	rating1, ok := myMap2["Varun"]
	if ok {
		println("Varun's rating is:", rating1)
	} else {
		println("Varun's age not found in map")
	}

	delete(myMap2, "Varun") // deleting an map key
	rating1, ok = myMap2["Varun"]
	if ok {
		println("Varun's rating is:", rating1)
	} else {
		println("Varun's age not found in map")
	}

	// Looping over map (or array or slice) using range
	for key, value := range myMap2 {
		fmt.Printf("key: %v | value: %v \n", key, value)
	}
}

func main() {
	// showcaseArrays()
	// showcaseSlices()
	// showcaseMaps()
	performanceTest()
}
