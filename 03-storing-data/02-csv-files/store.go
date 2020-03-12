package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// -- ----------------------------------
type Post struct {
	Id      int
	Content string
	Author  string
}

// -- ----------------------------------
func main() {

	// -- ------------------------------
	// Create a CSV file:

	csvFile, err := os.Create("posts.csv")

	if err != nil {
		panic(err)
	}

	defer csvFile.Close()

	// -- ------------------------------
	allPosts := []Post{
		Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
		Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
		Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}

	// -- ------------------------------
	// Write to CSV file:
	writer := csv.NewWriter(csvFile)

	for _, post := range allPosts {

		line := []string{
			strconv.Itoa(post.Id),
			post.Content,
			post.Author,
		}

		err := writer.Write(line)

		if err != nil {
			panic(err)
		}
	}

	// Flush writes any buffered data to the underlying io.Writer.
	writer.Flush()

	// -- ------------------------------
	// Read a CSV file:
	file, err := os.Open("posts.csv")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	// -- ------------------------------
	/*
		FieldsPerRecord
			If negative number:
				Indicates you aren’t that bothered if you don’t
				have all the fields in the record.

			If positive number:
				The number of fields you expect from each record
				and Go will throw an error if you get less.

			0:
				Use the number of fields in the first record as
				the FieldsPerRecord value.
	*/
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	record, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	// -- ------------------------------
	var posts []Post

	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
