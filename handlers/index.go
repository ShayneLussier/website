package handlers

import "github.com/labstack/echo"

func IndexHandler(ctx echo.Context) error {
	return ctx.File("templates/index.html")
}
