package handlers

import (
	"path/filepath"

	"github.com/labstack/echo"
)

func ResumePDFHandler(c echo.Context) error {
	fileName := "resume.pdf"
	filePath := "static/"

	return c.Attachment(filepath.Join(filePath, fileName), fileName)
}
