// handlers.go
package handlers

import (
	"github.com/labstack/echo"
)

// Register all Project Handlers
func RegisterProjectHandlers(e *echo.Echo, basePath string) {
	e.GET(basePath+"/tennis-streaming", TennisStreamingProjectHandler)
	e.GET(basePath+"/example-project", ExampleProjectHandler)
	// e.Get(basePath+"/new-project", NewProjectHandler)
}

func TennisStreamingProjectHandler(ctx echo.Context) error {
	return ctx.File("templates/html/projects/tennis-streaming.html")
}

func ExampleProjectHandler(ctx echo.Context) error {
	return ctx.File("templates/html/projects/example-project.html")
}
