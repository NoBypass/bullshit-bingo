package main

import (
	"bullshit-bingo/internal/controller"
	"bullshit-bingo/internal/db"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Renderer = &controller.Template{
		Templates: template.Must(template.ParseGlob("views/*.gohtml")),
	}

	conn := db.Connect()
	defer func() {
		if err := conn.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	gc := controller.NewGameController(conn)
	html := controller.NewHTMLController()

	e.GET("/join/:id", gc.Game)
	e.POST("/create", gc.Create)
	e.POST("/word", gc.CreateWord)
	e.DELETE("/word/:word", gc.DeleteWord)
	e.GET("/topics", gc.ListTopics)
	e.GET("/words", gc.ListWords)

	e.GET("/", html.Home)

	e.Start(":1234")
}
