package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Text interface {
	fetch(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}

type Post struct {
	Db      *sql.DB
	Id      int
	Content string
	Author  string
}

// Get a single post
func (post *Post) fetch(id int) (err error) {
	err = post.Db.QueryRow("select * from posts where id = ?", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

// Create a new post
func (post *Post) create() (err error) {
	stmt, err := post.Db.Prepare("insert into posts (content, author) values (?, ?)")
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
	_, err = post.Db.Exec("update posts set content = ?, author = ? where id = ?", post.Content, post.Author, post.Id)
	return
}

// Delete a post
func (post *Post) delete() (err error) {
	_, err = post.Db.Exec("delete from posts where id = ?", post.Id)
	return
}
