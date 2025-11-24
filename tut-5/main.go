package main

import "fmt"

func squareThisThing(thing2 [10]float32) [10]float32 {
	fmt.Printf("memory location of thing2 array: %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	fmt.Println("thing2 values: ", thing2)
	return thing2
}

func squareThisThingByRef(thing4 *[10]float32) [10]float32 {
	fmt.Printf("memory location of thing4 array: %p\n", thing4)
	for i := range thing4 {
		thing4[i] = thing4[i] * thing4[i]
	}
	fmt.Println("thing4 values: ", *thing4)
	return *thing4
}

func main() {
	myPointer := new(int)                                            // decalring a pointer and making it point to a int32 memory location
	fmt.Printf("The value myPointer points to is: %v\n", *myPointer) // dereferencing / fetching the value at the memory location the pointer is pointing to.
	// Notice the different uses of * here

	// assigning value to the memory location,
	// our pointer is pointing to:
	*myPointer = 10
	fmt.Printf("The value myPointer points to is: %v\n", *myPointer)

	// Pointer pointing to a existing variable
	myVar := 10
	myPointer = &myVar
	fmt.Printf("\n\nThe address myPointer stores is: %v\n", myPointer)
	fmt.Printf("The value myPointer points to is: %v\n", *myPointer)
	fmt.Printf("The value myVar points to is: %v\n", *myPointer)
	*myPointer = 100
	fmt.Printf("\n\nThe address myPointer stores is: %v\n", myPointer)
	fmt.Printf("The value myPointer points to is: %v\n", *myPointer)
	fmt.Printf("The value myVar points to is: %v\n", *myPointer)

	// Slices use arry of pointers under the hood, so they are exceptions to the variables copying values rule:
	originalSlice := []int{1, 2, 3}
	copySlice := originalSlice

	copySlice[2] = 4 // only modifying the copySlice
	fmt.Println("originalSlice: ", originalSlice)
	fmt.Println("copySlice: ", copySlice)

	// Usage of pointers in functions (sending by value VS reference)
	thing1 := [10]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("\n\nmemory location of thing1 array: %p\n", &thing1)
	result := squareThisThing(thing1)
	fmt.Println("thing1 values: ", thing1)
	fmt.Println("The result is: ", result)

	// sending by reference this time
	thing3 := [10]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("\n\nmemory location of thing3 array: %p\n", &thing3)
	result = squareThisThingByRef(&thing3)
	fmt.Println("thing3 values: ", thing3)
	fmt.Println("The result is: ", result)
}
