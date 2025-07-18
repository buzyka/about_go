package main

import (
	"github.com/buzyka/about_go/06_gin_project/internal/infrastructure/http/router"
)

func main() {
	r := router.SetupRouter() // Assuming you have a router package with SetupRouter function
	r.Run(":8080") // Start the server on port 8080
}
