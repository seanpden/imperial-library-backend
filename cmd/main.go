package main

import (
	"fmt"

	"github.com/seanpden/imperial-library-backend/internal/models"
	"github.com/seanpden/imperial-library-backend/internal/routes"
)

func main() {
	fmt.Println("Hello world!")

	r := routes.SetupRouter()
	r.Run()
}
