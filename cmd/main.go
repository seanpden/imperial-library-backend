package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/seanpden/imperial-library-backend/internal/routes"
)

func main() {
	// log.Println("Opening database...")
	// db, err := sql.Open("sqlite3", "./database.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Pinging database...")
	// pingerr := db.Ping()
	// if pingerr != nil {
	// 	log.Fatal(pingerr)
	// }
	// log.Println("Connected!")

	log.Println("Setting up router...")
	r := routes.SetupRouter()

	log.Println("Running application!")
	r.Run()
}
