package main

import (
	"fmt"

	"github.com/seanpden/imperial-library-backend/internal/routes"
)

func main() {
	fmt.Println("Hello world!")

	println("Setting up router...")
	r := routes.SetupRouter()

	println("Running...")
	r.Run()
}
