package handlers

import (
	"path/filepath"

	"github.com/labstack/echo"
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
	e.GET(basePath+"/example-project", ExampleProjectHandler)
	// e.Get(basePath+"/new-project", NewProjectHandler)
}

func IndexHandler(ctx echo.Context) error {
	return ctx.File("templates/html/index.html")
}

func ArchiveHandler(ctx echo.Context) error {
	return ctx.File("templates/html/archive.html")
}

func ResumePDFHandler(c echo.Context) error {
	fileName := "resume.pdf"
	filePath := "static/"

	return c.Attachment(filepath.Join(filePath, fileName), fileName)
}

func TennisStreamingProjectHandler(ctx echo.Context) error {
	return ctx.File("templates/html/projects/tennis-streaming.html")
}

func ExampleProjectHandler(ctx echo.Context) error {
	return ctx.File("templates/html/projects/example-project.html")
}
