package main

import (
	"log"
	"tasbih/pkg/database"
	"tasbih/pkg/handlers"
	"tasbih/pkg/store"

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

	db := database.InitDB()
	us := store.NewUserStore(db)
	h := handlers.NewHandler(us)

	//pages
	e.GET("/ping", handlers.Pong)
	e.GET("/", handlers.ServeLanding)
	e.GET("/signup", handlers.Signup_page)
	e.GET("/login", handlers.Login_page)
	e.GET("/home", handlers.Home_page)

	e.POST("/inc", handlers.Increment)
	e.GET("/ws", handlers.ChatWS)
	//auth
	e.POST("/auth/signup", h.CreateUser)
	e.POST("/auth/login", h.Login)
	e.POST("/auth/email", h.ValidateEmail)

	e.Debug = true

	e.Logger.Fatal(e.Start(":1323"))
}
