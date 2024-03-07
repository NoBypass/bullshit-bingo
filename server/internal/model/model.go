package model

type Word struct {
	Text    string `json:"text"`
	Checked bool   `json:"checked"`
}

type Game struct {
	ID    string
	Board [5][5]Word
	Msgs  map[string]chan string
}
