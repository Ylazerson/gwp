package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// -- --------------------------------------
func main() {

	// -- ----------------------------------
	data := []byte("Hello World!\n")

	// -- ----------------------------------
	//             APPROACH 1
	// Read/write using WriteFile and ReadFile
	err := ioutil.WriteFile("data1", data, 0644)

	if err != nil {
		panic(err)
	}

	read1, _ := ioutil.ReadFile("data1")

	fmt.Print(string(read1))

	// -- ----------------------------------
	//             APPROACH 2
	// write to file and read from file using the File struct
	file1, _ := os.Create("data2")

	/*
		A defer statement pushes a function call on a stack.
		The list of saved calls is then executed after the
		surrounding function returns.
	*/
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	// -- ----------------------------------
	file2, _ := os.Open("data2")
	defer file2.Close()

	// -- ----------------------------------
	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))
}