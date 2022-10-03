package models

import (
	"database/sql"
	"strconv"
)

type Tag struct {
	T_id     int    `json:"t_id"`
	Post_id  int    `json:"post_id"`
	Tag_data string `json:"tag_data"`
}

func GetTags(count int) ([]Tag, error) {
	rows, err := DB.Query("SELECT t_id,post_id,tag_data from tags LIMIT" + strconv.Itoa(count))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tags := make([]Tag, 0)
	for rows.Next() {
		singletag := Tag{}
		err = rows.Scan(&singletag.T_id, &singletag.Post_id, &singletag.Tag_data)
		if err != nil {
			return nil, err
		}
		tags = append(tags, singletag)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return tags, err
}

func GetTagById(id string) (Tag, error) {
	stmt, err := DB.Prepare("SELECT t_id,post_id,tag_data from tags where t_id=?")
	if err != nil {
		return Tag{}, err
	}
	tags := Tag{}
	sqlErr := stmt.QueryRow(id).Scan(&tags.T_id, &tags.Post_id, &tags.Tag_data)
	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Tag{}, err
		}
		return Tag{}, sqlErr
	}
	return tags, nil
}

func AddTags(newTag Tag) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}
	stmt, err := tx.Prepare("Insert into tags (post_id,tag_data) values (?,?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(newTag.Post_id, newTag.Tag_data)
	if err != nil {
		return false, err
	}
	tx.Commit()
	return true, nil
}

func UpdateTag(ourTag Tag, tagId int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}
	stmt, err := tx.Prepare("Update tags set tag_data=? where t_id=?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(ourTag.Tag_data, tagId)
	if err != nil {
		return false, err
	}
	tx.Commit()
	return true, nil
}

func DeleteTag(tagId int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}
	stmt, err := DB.Prepare("Delete from tags where t_id=?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(tagId)
	if err != nil {
		return false, err
	}
	tx.Commit()
	return true, nil
}
