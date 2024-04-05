package logic

import "github.com/google/uuid"

type Game struct {
	Board   *Board
	ID      string
	Players []string
}

func NewGame() *Game {
	id := uuid.New().String()
	return &Game{
		ID:      id,
		Board:   NewBoard(5),
		Players: make([]string, 0),
	}
}
