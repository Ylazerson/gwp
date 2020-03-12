package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

// -- ------------------------------------------
// Remember, the Comments field is a slice, which
// is really a pointer to an array.
type Post struct {
	Id       int
	Content  string
	Author   string
	Comments []Comment
}

type Comment struct {
	Id      int
	Content string
	Author  string
	Post    *Post
}

// -- ------------------------------------------
var Db *sql.DB

// -- ------------------------------------------
// Connect to the database:
func init() {

	var err error

	Db, err = sql.Open(
		"postgres",
		"user=gwp dbname=gwp password=gwp sslmode=disable",
	)

	if err != nil {
		panic(err)
	}
}

// -- ------------------------------------------
// Create comment on a post:
func (comment *Comment) Create() (err error) {

	if comment.Post == nil {
		err = errors.New("Post not found")
		return
	}

	var sqlStr string = `
	insert into comments 
	       (
			content, 
			author, 
			post_id) 
	values (
			$1, 
			$2, 
			$3) 
	returning id
	`

	err = Db.QueryRow(
		sqlStr,
		comment.Content,
		comment.Author,
		comment.Post.Id,
	).Scan(&comment.Id)

	return
}

// -- ------------------------------------------
// Get a single post
func GetPost(id int) (post Post, err error) {

	var sqlStr string

	// -- --------------------------------------
	post = Post{}
	post.Comments = []Comment{}

	// -- --------------------------------------
	sqlStr = `
	select   id, 
			 content, 
			 author 
	from     posts 
	where    id = $1
	`

	err = Db.QueryRow(
		sqlStr,
		id,
	).Scan(
		&post.Id,
		&post.Content,
		&post.Author,
	)

	// -- --------------------------------------
	sqlStr = `
	select   id, 
			 content, 
			 author 
	from     comments 
	where    post_id = $1
	`

	rows, err := Db.Query(
		sqlStr,
		id,
	)

	if err != nil {
		return
	}

	for rows.Next() {
		comment := Comment{Post: &post}

		err = rows.Scan(
			&comment.Id,
			&comment.Content,
			&comment.Author,
		)

		if err != nil {
			return
		}

		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}

// -- ------------------------------------------
// Create a new post
func (post *Post) Create() (err error) {

	var sqlStr string = `
	insert into posts 
		   (content, author) 
	values ($1, $2) 
	returning id
	`

	err = Db.QueryRow(
		sqlStr,
		post.Content,
		post.Author,
	).Scan(&post.Id)

	return
}

// -- ------------------------------------------
func main() {

	// -- --------------------------------------
	// NOTE, we're ignoring errors here just for brevity.

	// -- --------------------------------------
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	post.Create()

	// -- --------------------------------------
	// Add a comment
	comment := Comment{Content: "Good post!", Author: "Joe", Post: &post}
	comment.Create()
	readPost, _ := GetPost(post.Id)

	// -- --------------------------------------
	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0].Post)
}
