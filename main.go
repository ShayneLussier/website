package main

import (
	"website/handlers"

	"github.com/labstack/echo"
)

func main() {
	router := echo.New()

	// Register Routes
	handlers.RegisterRoutes(router)
	// Register Project Routes
	handlers.RegisterProjectRoutes(router, "/projects")

	router.Logger.Fatal(router.Start("localhost:3000"))
}
