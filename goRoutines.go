package main

import (
	"fmt"
	"sync"
)

var counter = 0
var wg = sync.WaitGroup{}

func printCount() {
	fmt.Println(counter)
	wg.Done() // Removes one group from the wait group
}
func incrementCounter() {
	counter++
	wg.Done()
}
func main() {
	// Go routines helps in concurrency. They helps us to run multiple functions at the same time parallely (similar to threads)
	// just write go keyword in front of the func to trigger a go routines
	for i := 1; i <= 4; i++ {
		wg.Add(1)             // Adds one group to the wait group
		go incrementCounter() // New routine will be created
		wg.Wait()             // Waits until all the groups in the wait group gets completed
		wg.Add(1)
		go printCount() // New routine will be created
		wg.Wait()
	}
	// Mutex can also be used to synchronize and avoid the race conditions
	// Race conditions occur when different processes are reading and writing at the same memory - this leads to data inconsistency
	// If no wait group is used in this example then the result will not be as expected
}
