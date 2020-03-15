package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// -- -------------------------------------
type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// -- -------------------------------------
func main() {

	// -- ---------------------------------
	// Open JSON file:
	jsonFile, err := os.Open("post.json")

	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}

	defer jsonFile.Close()

	// -- ---------------------------------
	// Create JSON Decoder
	decoder := json.NewDecoder(jsonFile)

	// -- ---------------------------------
	// Decode JSON into Struct
	//
	// Iterate until EOF is detected:
	for {
		var post Post

		err := decoder.Decode(&post)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		fmt.Println(post)
	}
}
