package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	dbData = [5]string{"id1", "id2", "id3", "id4", "id5"}
	wg     = sync.WaitGroup{} // waitGroup init
)

func dbCall(i int) {
	// Simulator for a dbCall()
	delay := rand.Float32() * 2000 // times 2000 means the delay will be anywhere between 0-2 seconds
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("The results from the database are: %v\n", dbData[i])
}

func dbConcurrentCall(i int) {
	// Simulator for a dbCall()
	delay := rand.Float32() * 2000 // times 2000 means the delay will be anywhere between 0-2 seconds
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("The results from the database are: %v\n", dbData[i])
	wg.Done()
}

func sequenctialDBcalls() {
	t0 := time.Now()
	for i := range dbData {
		dbCall(i) // making 5 db calls in sequence
	}
	fmt.Printf("\nThe total execution time: %v\n\n", time.Since(t0))
}

func concurrentDBcalls() {
	t0 := time.Now()
	for i := range dbData {
		wg.Add(1)
		go dbConcurrentCall(i) // go keyword makes func concurrrent
	}
	wg.Wait()
	fmt.Printf("\nThe total execution time: %v\n\n", time.Since(t0))
}

func main() {
	sequenctialDBcalls()
	concurrentDBcalls()
}
