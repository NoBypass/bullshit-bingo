package main

import (
	"bullshit-bingo/internal/controller"
	"bullshit-bingo/internal/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	conn := db.Connect(os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASS"))
	gc := controller.NewGameController(conn)

	e.GET("/join/:id", gc.Game)
	e.POST("/create/:id", gc.Create)

	e.Start(":1234")
}
