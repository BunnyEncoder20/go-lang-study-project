package main

import (
	"fmt"
	"time"
)

func performanceTest() {
	n := 1000000
	unallocTestSlice := []int{}
	allocTestSlice := make([]int, 0, n)
	fmt.Println("Total time with  unallocated memory: ", timeLoop(unallocTestSlice, n))
	fmt.Println("Total time with preallocated memory: ", timeLoop(allocTestSlice, n))
}

func timeLoop(testSlice []int, n int) time.Duration {
	t0 := time.Now()
	for i := range n {
		testSlice = append(testSlice, i)
	}
	return time.Since(t0)
}
