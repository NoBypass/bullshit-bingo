package controller

import (
	"bullshit-bingo/internal/model"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/websocket"
	"net/http"
	"strconv"
	"strings"
)

type GameCache struct {
	Games map[string]model.Game
	DB    *mongo.Client
}

func NewGameController(db *mongo.Client) *GameCache {
	return &GameCache{
		Games: make(map[string]model.Game),
		DB:    db,
	}
}

func (gc *GameCache) Create(c echo.Context) error {
	gameID := c.Param("id")
	if _, ok := gc.Games[gameID]; ok {
		return c.String(http.StatusConflict, fmt.Sprintf("game with id %s already exists", gameID))
	}
	game := model.Game{
		ID:   gameID,
		Msgs: make(map[string]chan string),
	}
	gc.Games[gameID] = game
	return c.String(http.StatusCreated, fmt.Sprintf("game with id %s created", gameID))
}

func (gc *GameCache) Game(c echo.Context) error {
	gameID := c.Param("id")
	game, ok := gc.Games[gameID]
	if !ok {
		return c.String(http.StatusNotFound, fmt.Sprintf("couldn't find game with id %s", gameID))
	}
	receiver := make(chan string)
	defer close(receiver)
	playerUUID := uuid.New().String()
	game.Msgs[playerUUID] = receiver
	defer delete(game.Msgs, playerUUID)

	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		data, err := json.Marshal(game.Board)
		if err != nil {
			c.Logger().Error(err)
		}

		err = websocket.Message.Send(ws, data)
		if err != nil {
			c.Logger().Error(err)
		}

		client := make(chan string)
		defer close(client)
		go func() {
			for {
				msg := ""
				err = websocket.Message.Receive(ws, &msg)
				if err != nil {
					c.Logger().Error(err)
				}
				client <- msg
			}
		}()

		for {
			select {
			case msg := <-receiver:
				err = websocket.Message.Send(ws, msg)
				if err != nil {
					c.Logger().Error(err)
				}
			case msg := <-client:
				for _, ch := range game.Msgs {
					ch <- msg
				}

				split := strings.Split(msg, ":")
				if len(split) != 2 {
					continue
				}

				var (
					enable = split[0] == "enable"
					idStr  = split[1]
				)

				id, err := strconv.Atoi(idStr)
				if err != nil {
					c.Logger().Error(err)
					continue
				}

				for i, row := range game.Board {
					if id > (i+1)*len(game.Board) {
						continue
					}
					for j, word := range row {
						if id == (j+1)*(i+1) {
							word.Enabled = enable
						}
					}
				}
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
