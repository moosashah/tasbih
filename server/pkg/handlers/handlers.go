package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type header struct {
	Title string
}

type page struct {
	header
}

type pong struct {
	Message string `json:"message"`
}

func Pong(c echo.Context) error {
	return c.JSON(http.StatusOK, pong{
		Message: "pong",
	})
}

var count int

func Increment(c echo.Context) error {
	count++
	return c.String(http.StatusOK, strconv.Itoa(count))
}
