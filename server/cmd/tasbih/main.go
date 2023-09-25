package main

import (
	"tasbih/pkg/database"
	"tasbih/pkg/handlers"
	"tasbih/pkg/store"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	db := database.InitDB()
	us := store.NewUserStore(db)
	h := handlers.NewHandler(us)

	//pages
	e.GET("/ping", handlers.Pong)

	e.POST("/inc", handlers.Increment)

	//auth
	e.POST("/auth/signup", h.CreateUser)
	e.POST("/auth/login", h.Login)

	e.Debug = true

	e.Logger.Fatal(e.Start(":1323"))
}
