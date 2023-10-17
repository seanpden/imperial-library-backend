package models

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Read all book info
// Read a specific book info
// db, err = sql.Open("sqlite3", "./database.db")

func GetAllBookInfos() ([]BookInfo, error) {
	db := OpenDatabase()
	var book_infos []BookInfo

	query := `
  SELECT
    title,
    COALESCE(link_to_content, '') as author,
    COALESCE(author, '') as author,
    COALESCE(summary, '') as summary
  FROM 
    book_info
  `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// loop throuhg rows, using Scan to assign column data to struct fields
	for rows.Next() {
		var book_info BookInfo
		if err := rows.Scan(&book_info.Title, &book_info.Link_To_Content, &book_info.Author, &book_info.Summary); err != nil {
			return nil, fmt.Errorf("GetAllBookInfos: %v", err)
		}
		book_infos = append(book_infos, book_info)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAllBookInfos: %v", err)
	}

	return book_infos, nil
}
