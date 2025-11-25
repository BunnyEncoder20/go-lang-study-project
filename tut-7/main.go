package main

import (
	"fmt"
	"time"
)

func main() {
	// channel := make(chan int) // make + chan keywords
	// go process(channel)
	// for i := range channel {
	// 	fmt.Println(i)
	// }

	bufferChannel := make(chan int, 5)
	go process(bufferChannel)
	for i := range bufferChannel {
		time.Sleep(1 * time.Second) // some work main is doing
		fmt.Println(i)
	}
}

func process(channel chan int) {
	defer close(channel)
	for i := range 5 {
		channel <- i
	}
	fmt.Println("Exiting process...")
}
