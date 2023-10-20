package models

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db = OpenDatabase()

func GetAllBooks() ([]Book, error) {
	var books []Book

	query := `
  SELECT
    TRIM(i.title) as [i.title],
    COALESCE(c.author, '') as [c.author], -- cleaner
    TRIM(COALESCE(i.summary, '')) as [i.summary],
    c.link_to_content,
    COALESCE(c.comment, '') as [c.comment],
    c.text
  FROM 
    book_content c,
    book_info i
  WHERE 
    i.link_to_content = c.link_to_content
  `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// loop throuhg rows, using Scan to assign column data to struct fields
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.Title, &book.Author, &book.Summary, &book.Link_To_Content, &book.Comment, &book.Text); err != nil {
			return nil, fmt.Errorf("GetAllBookContents: %v", err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAllBookContents: %v", err)
	}

	return books, nil
}

func GetBooksFuzzy(titleq string, authorq string, textq string) ([]Book, error) {
	var books []Book

	query := `
  SELECT
    TRIM(i.title) as [i.title],
    COALESCE(c.author, '') as [c.author], -- cleaner
    TRIM(COALESCE(i.summary, '')) as [i.summary],
    c.link_to_content,
    COALESCE(c.comment, '') as [c.comment],
    c.text
  FROM 
    book_content c,
    book_info i
  WHERE 
    i.link_to_content = c.link_to_content
    and i.title like ?
    and c.author like ?
    and c.text like ?
    
  `

	titleq = fmt.Sprintf("%%%s%%", titleq)
	authorq = fmt.Sprintf("%%%s%%", authorq)
	textq = fmt.Sprintf("%%%s%%", textq)
	rows, err := db.Query(query, titleq, authorq, textq)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.Title, &book.Author, &book.Summary, &book.Link_To_Content, &book.Comment, &book.Text); err != nil {
			return nil, fmt.Errorf("GetBookInfoFuzzy: %v", err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetBookInfoFuzzy: %v", err)
	}

	return books, nil
}
