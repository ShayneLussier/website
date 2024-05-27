package main

import (
	"website/handlers"

	"github.com/labstack/echo"
)

func main() {
	router := echo.New()

	router.GET("/", handlers.IndexHandler)

	router.Logger.Fatal(router.Start("localhost:3000"))
}
