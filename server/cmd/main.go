package main

import (
	"bullshit-bingo/internal/controller"
	"bullshit-bingo/internal/db"
	"context"
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
	defer func() {
		if err := conn.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	gc := controller.NewGameController(conn)

	e.GET("/join/:id", gc.Game)
	e.POST("/create", gc.Create)
	e.POST("/word", gc.CreateWord)
	e.DELETE("/word/:word", gc.DeleteWord)
	e.GET("/topics", gc.ListTopics)
	e.GET("/words", gc.ListWords)

	e.Start(":1234")
}
