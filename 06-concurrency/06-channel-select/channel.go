// B''H

package main

import (
	"fmt"
)

// -- ------------------------------------
func callerA(c chan string) {
	c <- "Hello World!"
	close(c)
}

func callerB(c chan string) {
	c <- "Hola Mundo!"
	close(c)
}

// -- ------------------------------------
func main() {
	a, b := make(chan string), make(chan string)

	go callerA(a)
	go callerB(b)

	// -- --------------------------------
	// Loop thru open channels:

	var msg string
	openA, openB := true, true

	for openA || openB {

		// In each iteration the Go runtime determines
		// whether you receive from channel a or channel b,
		// depending on the channel that has a value at
		// the time of selection.
		// If both are available, the Go runtime will randomly pick one.
		select {
		case msg, openA = <-a:
			if openA {
				fmt.Printf("%s from A\n", msg)
			}
		case msg, openB = <-b:
			if openB {
				fmt.Printf("%s from B\n", msg)
			}
		}
	}
}
