package handlers

import (
	"path/filepath"

	"github.com/labstack/echo/v4"
)

// Register all routes
func RegisterRoutes(e *echo.Echo) {
	e.GET("/", IndexHandler)
	e.GET("/archive", ArchiveHandler)
	e.GET("/resume.pdf", ResumePDFHandler)
	// e.Get("/path", HandlerName)
}

func RegisterProjectRoutes(e *echo.Echo, basePath string) {
	e.GET(basePath+"/tennis-streaming", TennisStreamingProjectHandler)
	// e.Get(basePath+"/new-project", NewProjectHandler)
	e.GET(basePath+"/example-project", ExampleProjectHandler)
}

func IndexHandler(ctx echo.Context) error {
	return ctx.File("templates/index.html")
}

func ArchiveHandler(ctx echo.Context) error {
	return ctx.File("templates/archive.html")
}

func ResumePDFHandler(c echo.Context) error {
	fileName := "resume.pdf"
	filePath := "static/"

	return c.Attachment(filepath.Join(filePath, fileName), fileName)
}

func TennisStreamingProjectHandler(ctx echo.Context) error {
	return ctx.File("templates/projects/tennis/tennis-streaming.html")
}

// Example
func ExampleProjectHandler(ctx echo.Context) error {
	return ctx.File("templates/projects/project-example/example-project.html")
}
