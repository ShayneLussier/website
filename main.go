package main

import (
	"website/handlers"

	"github.com/labstack/echo"
)

func main() {
	router := echo.New()

	// Main and sub pages
	router.GET("/", handlers.IndexHandler)
	router.GET("/archive", handlers.ArchiveHandler)

	//	PDF
	router.GET("/resume.pdf", handlers.ResumePDFHandler)

	// Projects
	handlers.RegisterProjectHandlers(router, "/projects")

	router.Logger.Fatal(router.Start("localhost:3000"))
}
