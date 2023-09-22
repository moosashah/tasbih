package main

import (
	"log"
	"tasbih/pkg/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	tmpl, err := handlers.ParseHtmlFiles("./views")
	if err != nil {
		log.Fatalf("could not initialize templates: %+v", err)
	}
	e.Renderer = handlers.NewTemplateRenderer(tmpl)

	e.Static("/scripts", "scripts")
	e.Use(middleware.Logger())

	e.GET("/ping", handlers.Pong)
	e.GET("/", handlers.ServeLanding)
	e.POST("/inc", handlers.Increment)
	e.GET("/ws", handlers.ChatWS)

	e.Logger.Fatal(e.Start(":1323"))
}
