package models

import (
	"strconv"
)

type Author struct {
	Author_Name string `json:"author_name"`
}

func GetAuthor(count int) ([]Author, error) {
	rows, err := DB.Query("SELECT author_name from posts LIMIT" + strconv.Itoa(count))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	author := make([]Author, 0)
	for rows.Next() {
		singleauthor := Author{}
		err = rows.Scan(&singleauthor.Author_Name)
		if err != nil {
			return nil, err
		}
		author = append(author, singleauthor)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return author, err
}
