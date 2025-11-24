package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	// The results slice must be protected by a Mutex.
	results []int
	dbData  = []int{42, 99, 10, 75, 51}
	wg      sync.WaitGroup
	mu      sync.Mutex // The Mutex to guard access to the 'results' slice
	rwmu    sync.RWMutex
)

func dbCall(i int) {
	// Simulator for a dbCall()
	delay := rand.Float32() * 2000 // times 2000 means the delay will be anywhere between 0-2 seconds
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("The results from the database are: %v\n", dbData[i])
}

func sequenctialDBcalls() {
	t0 := time.Now()
	for i := range dbData {
		dbCall(i) // making 5 db calls in sequence
	}
	fmt.Printf("\nThe total execution time: %v\n\n", time.Since(t0))
}

func dbConcurrentCall(i int) {
	// Simulator for a dbCall()
	delay := rand.Float32() * 2000 // times 2000 means the delay will be anywhere between 0-2 seconds
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("The results from the database are: %v\n", dbData[i])
	mu.Lock() // lock the data struct before making edits
	results = append(results, dbData[i])
	mu.Unlock() // Unlock the data struct afterwards

	wg.Done()
}

func concurrentDBcalls() {
	t0 := time.Now()
	for i := range dbData {
		wg.Add(1)
		go dbConcurrentCall(i) // go keyword makes func concurrrent
	}
	wg.Wait()
	fmt.Printf("\nThe total execution time: %v\n\n", time.Since(t0))
	fmt.Printf("The Results are %v\n\n", results)
}

func main() {
	sequenctialDBcalls()
	concurrentDBcalls()
}
