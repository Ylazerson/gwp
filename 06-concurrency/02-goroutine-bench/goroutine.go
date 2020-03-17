package main

// import "fmt"
import "time"

// -- -------------------------------
/*
Benchmark with 2 CPU's:
go test -run x -bench . -cpu 2

Benchmark with 4 CPU's:
go test -run x -bench . -cpu 4

Benchmark with all available CPU's:
go test -run x -bench .
*/
// -- -------------------------------

func numsNoSleep() {
	for i := 0; i < 100; i++ {
		// fmt.Printf("%d ", i)
	}
}

func lettersNoSleep() {
	for i := 'A'; i < 'A'+100; i++ {
		// fmt.Printf("%c ", i)
	}
}

func numsSleep() {
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Microsecond)
		// fmt.Printf("%d ", i)
	}
}

func lettersSleep() {
	for i := 'A'; i < 'A'+100; i++ {
		time.Sleep(1 * time.Microsecond)
		// fmt.Printf("%c ", i)
	}
}

func printNoSleep() {
	numsNoSleep()
	lettersNoSleep()
}

func goPrintNoSleep() {
	go numsNoSleep()
	go lettersNoSleep()
}

func printSleep() {
	numsSleep()
	lettersSleep()
}

func goPrintSleep() {
	go numsSleep()
	go lettersSleep()
}

func main() {
}
