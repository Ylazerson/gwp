// B''H

package main

import (
	"fmt"
	"time"
)

// -- -------------------------------------
func nums() {

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}

}

// -- -------------------------------------
func letters() {

	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}

}

// -- -------------------------------------
func main() {

	// -- ---------------------------------
	fmt.Println("\n Kicking off goroutines - round 1")
	go nums()
	go letters()

	// Wait so we can get output from goroutines.
	time.Sleep(1 * time.Millisecond)

	// -- ---------------------------------
	fmt.Println("\n Kicking off goroutines - round 2")
	// Below won't show any output because the program ends
	// before the goroutines could output anything :)
	go nums()
	go letters()

}
