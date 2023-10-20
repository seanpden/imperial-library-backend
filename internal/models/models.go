package models

import (
	"database/sql"
	"log"
)

type Book struct {
	Title           string `json:"title"`
	Author          string `json:"author"`
	Summary         string `json:"summary"`
	Link_To_Content string `json:"link_to_content"`
	Comment         string `json:"comment"`
	Text            string `json:"text"`
}

type PGDBConfig struct {
	DATABASE_URL string
	PGDATABASE   string
	PGHOST       string
	PGPASSWORD   string
	PGPORT       string
	PGUSER       string
}

func (pg *PGDBConfig) ConstructCNXNSTR() string {
	return "./internal/models/database.db"
}

func ConstructDBConnection() string {
	dbc := PGDBConfig{}
	cnxnstr := dbc.ConstructCNXNSTR()
	return cnxnstr
}

func OpenDatabase() *sql.DB {
	log.Println("Opening database...")
	cnxnstr := ConstructDBConnection()
	db, err := sql.Open("sqlite3", cnxnstr)
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
