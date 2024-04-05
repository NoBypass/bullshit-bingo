package controller

import (
	"bullshit-bingo/internal/logic"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HTMLController struct {
	Games map[string]*logic.Game
}

func NewHTMLController() *HTMLController {
	return &HTMLController{
		Games: make(map[string]*logic.Game),
	}
}

func (hc *HTMLController) Home(c echo.Context) error {
	gameID := c.QueryParam("game")
	var game *logic.Game
	if gameID != "" {
		g, ok := hc.Games[gameID]
		if !ok {
			return c.Render(http.StatusNotFound, "error.gohtml", nil)
		}
		game = g
	}

	if game == nil {
		game = logic.NewGame()
		hc.Games[game.ID] = game
	}
	return c.Render(http.StatusOK, "index.gohtml", game)
}
