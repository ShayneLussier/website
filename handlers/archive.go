package handlers

import "github.com/labstack/echo"

func ArchiveHandler(ctx echo.Context) error {
	return ctx.File("templates/archive.html")
}
