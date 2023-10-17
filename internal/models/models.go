package models

import (
	"database/sql"
	"log"
)

type BookContent struct {
	Link_To_Content string `json:"link_to_content"`
	Author          string `json:"author"`
	Comment         string `json:"comment"`
	Text            string `json:"text"`
}

func (b *BookContent) TableName() string {
	return "book_content"
}

type BookInfo struct {
	Title           string `json:"title"`
	Link_To_Content string `json:"link_to_content"`
	Author          string `json:"author"`
	Summary         string `json:"summary"`
}

func (b *BookInfo) TableName() string {
	return "book_info"
}

type PGDBConfig struct {
	DATABASE_URL string
	PGDATABASE   string
	PGHOST       string
	PGPASSWORD   string
	PGPORT       string
	PGUSER       string
}

type ConnectionString struct {
	CONNECTIONSTRING string
}

func OpenDatabase() *sql.DB {
	log.Println("Opening database...")
	db, err := sql.Open("sqlite3", "./internal/models/database.db")
	// db, err := sql.Open("sqlite3", "../../cmd/database.db")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Pinging database...")
	pingerr := db.Ping()
	if pingerr != nil {
		log.Panic(pingerr)
	}
	log.Println("Connected!")

	return db
}
