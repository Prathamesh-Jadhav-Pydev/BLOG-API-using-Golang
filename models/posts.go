package models

import (
	"database/sql"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./blog.db")
	if err != nil {
		return err
	}
	DB = db
	return nil
}

type Posts struct {
	Post_id     int    `json:"post_id"`
	Author_Name string `json:"author_name"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}

func GetPosts(count int) ([]Posts, error) {
	rows, err := DB.Query("SELECT post_id,author_name,title,content from posts LIMIT" + strconv.Itoa(count))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := make([]Posts, 0)
	for rows.Next() {
		singlepost := Posts{}
		err = rows.Scan(&singlepost.Post_id, &singlepost.Author_Name, &singlepost.Title, &singlepost.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, singlepost)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return posts, err
}
func GetPersonById(id string) (Posts, error) {
	stmt, err := DB.Prepare("SELECT post_id,author_name,title,content from posts where post_id=?")
	if err != nil {
		return Posts{}, err
	}
	posts := Posts{}
	sqlErr := stmt.QueryRow(id).Scan(&posts.Post_id, &posts.Author_Name, &posts.Title, &posts.Content)
	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Posts{}, err
		}
		return Posts{}, sqlErr
	}
	return posts, nil
}
func AddPosts(newPost Posts) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}
	stmt, err := tx.Prepare("Insert into posts (author_name,title,content) values (?,?,?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	//_, err = stmt.Exec(newPost.Author_Name, newPost.Title, newPost.Content)
	_, err = stmt.Exec(newPost.Author_Name, newPost.Title, newPost.Content)
	if err != nil {
		return false, err
	}
	tx.Commit()
	return true, nil
}

func UpdatePost(ourPosts Posts, postId int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}
	stmt, err := tx.Prepare("Update posts set author_name=?, title=?, content=? where post_id=?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(ourPosts.Author_Name, ourPosts.Title, ourPosts.Content, postId)
	if err != nil {
		return false, err
	}
	tx.Commit()
	return true, nil
}

func DeletePost(postId int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}
	stmt, err := DB.Prepare("Delete from posts where post_id=?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(postId)
	if err != nil {
		return false, err
	}
	tx.Commit()
	return true, nil
}
