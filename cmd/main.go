package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/seanpden/imperial-library-backend/internal/routes"
)

func main() {
	log.Println("Setting up router...")
	r := routes.SetupRouter()

	log.Println("Running application!")
	r.Run()
}
