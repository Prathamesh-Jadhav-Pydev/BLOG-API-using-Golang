package models

import (
	"database/sql"
	"strconv"
)

type Comment struct {
	Comment_id   int    `json:"c_id"`
	Post_id      int    `json:"post_id"`
	Comment_data string `json:"comment_data"`
}

func GetComment(count int) ([]Comment, error) {
	rows, err := DB.Query("SELECT c_id,post_id,comment_data from comments LIMIT" + strconv.Itoa(count))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comment := make([]Comment, 0)
	for rows.Next() {
		singlecomment := Comment{}
		err = rows.Scan(&singlecomment.Comment_id, &singlecomment.Post_id, &singlecomment.Comment_data)
		if err != nil {
			return nil, err
		}
		comment = append(comment, singlecomment)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return comment, err
}

func GetCommentById(id string) (Comment, error) {
	stmt, err := DB.Prepare("SELECT c_id,post_id,comment_data from comments where c_id=?")
	if err != nil {
		return Comment{}, err
	}
	comment := Comment{}
	sqlErr := stmt.QueryRow(id).Scan(&comment.Comment_id, &comment.Post_id, &comment.Comment_data)
	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Comment{}, err
		}
		return Comment{}, sqlErr
	}
	return comment, nil
}

func AddComment(newComment Comment) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}
	stmt, err := tx.Prepare("Insert into comments (post_id,comment_data) values (?,?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(newComment.Post_id, newComment.Comment_data)
	if err != nil {
		return false, err
	}
	tx.Commit()
	return true, nil
}

func UpdateComment(ourComment Comment, commentId int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}
	stmt, err := tx.Prepare("Update comments set comment_data=? where c_id=?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(ourComment.Comment_data, commentId)
	if err != nil {
		return false, err
	}
	tx.Commit()
	return true, nil
}

func DeleteComment(commentId int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}
	stmt, err := DB.Prepare("Delete from comments where c_id=?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(commentId)
	if err != nil {
		return false, err
	}
	tx.Commit()
	return true, nil
}
