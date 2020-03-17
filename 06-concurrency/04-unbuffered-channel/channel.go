package main

import (
	"fmt"
	"time"
)

/*
// -- ---------------------------------------
Note sometimes the Caught statement is printed before
the Threw, but that’s not important - it’s just the
runtime scheduling between the print statements
after sending or receiving from the channel.

What's more important to notice is that the numbers are
in sequence, meaning once the number is “thrown” from
the thrower, the same number must be “caught” by
the catcher before proceeding to the next number.
// -- ---------------------------------------
*/

func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		// Pass number into channel:
		fmt.Println("Threw  >>", i)
	}
}

func catcher(c chan int) {
	for i := 0; i < 5; i++ {
		num := <-c
		// Take number from channel:
		fmt.Println("Caught <<", num)
	}
}

func main() {
	c := make(chan int)
	go thrower(c)
	go catcher(c)
	time.Sleep(100 * time.Millisecond)
}
