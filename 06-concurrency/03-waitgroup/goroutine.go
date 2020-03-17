// B''H

package main

import (
	"fmt"
	"sync"
	"time"
)

// -- ------------------------------------
func nums(wg *sync.WaitGroup) {

	for i := 0; i < 10; i++ {

		// Sleep to simulate "work"
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}

	// Decrement counter:
	wg.Done()
}

// -- ------------------------------------
func letters(wg *sync.WaitGroup) {

	for i := 'A'; i < 'A'+10; i++ {

		// Sleep to simulate "work"
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}

	// Decrement counter:
	wg.Done()
}

// -- ------------------------------------
func main() {

	// Declare WaitGroup:
	var wg sync.WaitGroup

	// Set up counter:
	wg.Add(2)

	go nums(&wg)
	go letters(&wg)

	// Block until counter reaches 0:
	wg.Wait()
}
