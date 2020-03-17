// B''H

package main

/*
// -- -----------------------------------
SAMPLE OUTPUT:

BenchmarkPrintNoSleep-4         20000000               103 ns/op
BenchmarkGoPrintNoSleep-4        3000000               531 ns/op
BenchmarkPrintSleep-4                100          13311274 ns/op
BenchmarkGoPrintSleep-4           200000             11088 ns/op

NOTE:
    BenchmarkGoPrintSleep is about 1,200 times quicker!
    (13311274/11088) = 1200.5

NOTE:
- BenchmarkGoPrintNoSleep is slower than BenchmarkPrintNoSleep.
- Remember, there’s no such thing as a free lunch.
- Starting up goroutines has a cost no matter how lightweight it is.
- The functions here (with sleep) are so trivial and it ran so quickly
  that the costs of using goroutines outweigh those of running them sequentially.
// -- -----------------------------------
*/

import "testing"

// import "time"

// normal run
func BenchmarkPrintNoSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printNoSleep()
	}
}

// run with goroutines
func BenchmarkGoPrintNoSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrintNoSleep()
	}
}

// run with some work
func BenchmarkPrintSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printSleep()
	}
}

// run with goroutines and some work
func BenchmarkGoPrintSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrintSleep()
	}
}
