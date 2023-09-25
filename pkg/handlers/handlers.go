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

func Pong(c echo.Context) error {
	return c.Render(http.StatusOK, "ping.html", page{
		header{
			Title: "Test page",
		},
	})
}

func ServeLanding(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", page{
		header{
			Title: "Tasbih",
		},
	})
}

var count int

func Increment(c echo.Context) error {
	count++
	return c.String(http.StatusOK, strconv.Itoa(count))
}
