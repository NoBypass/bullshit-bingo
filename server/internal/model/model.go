package model

type Word struct {
	Text    string `json:"text"`
	Enabled bool   `json:"enabled"`
}

type Game struct {
	ID    string
	Board [5][5]Word
	Msgs  map[string]chan string
}
