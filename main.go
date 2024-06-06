package main

import (
	"website/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()

	// Define static file directory
	router.Static("/", "static")

	// Register Routes
	handlers.RegisterRoutes(router)
	// Register Project Routes
	handlers.RegisterProjectRoutes(router, "/projects")
	//Register Script Routes
	handlers.RegisterScripts(router)

	router.Logger.Fatal(router.Start("localhost:3000"))
}
