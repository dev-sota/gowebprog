package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

// connect to the Db
func init() {
	var err error
	Db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/gwp")
	if err != nil {
		panic(err)
	}
}

// Get a single post
func retrieve(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select * from posts where id = ?", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

// Create a new post
func (post *Post) create() (err error) {
	stmt, err := Db.Prepare("insert into posts (content, author) values (?, ?)")
	if err != nil {
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(post.Content, post.Author)
	postId, err := res.LastInsertId()
	post.Id = int(postId)
	return
}

// Update a post
func (post *Post) update() (err error) {
	_, err = Db.Exec("update posts set content = ?, author = ? where id = ?", post.Content, post.Author, post.Id)
	return
}

// Delete a post
func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id = ?", post.Id)
	return
}
