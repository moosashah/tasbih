package main

import (
	"log"
	"tasbih/pkg/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	tmpl, err := handlers.ParseHtmlFiles("./views")
	if err != nil {
		log.Fatalf("could not initialize templates: %+v", err)
	}
	e.Renderer = handlers.NewTemplateRenderer(tmpl)

	e.GET("/ping", handlers.Pong)
	e.GET("/", handlers.ServeLanding)
	e.POST("/inc", handlers.Increment)

	e.Logger.Fatal(e.Start(":1323"))
}
