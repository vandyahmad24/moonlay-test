package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/vandyahmad24/moonlay-test/app/helper"
	"net/http"
	"path/filepath"
)

// FileTypeValidator is a middleware that validates the file type
func FileTypeValidator(allowedTypes []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get the uploaded file
			file, err := c.FormFile("file")
			if err != nil {
				fmt.Println("file kosong")
				return next(c)
			}
			// Get the file extension
			ext := filepath.Ext(file.Filename)
			// Check if the file extension is in the list of allowed types
			valid := false
			for _, t := range allowedTypes {
				if t == ext {
					valid = true
					break
				}
			}
			if !valid {
				response := helper.ApiWithOutData(false, "Invalid file type")
				return c.JSON(http.StatusBadRequest, response)
			}
			return next(c)
		}
	}
}
