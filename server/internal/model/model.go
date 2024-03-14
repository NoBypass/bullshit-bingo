package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Word struct {
	Text    string `json:"text"`
	Checked bool   `json:"checked"`
}

type Game struct {
	ID    string
	Board [5][5]Word
	Msgs  map[string]chan string
}

type WordRequest struct {
	Topics []string `json:"topics"`
	Text   string   `json:"text"`
}

type DBWord struct {
	ID string `json:"_id"`
	WordRequest
}

type Topic struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Text string             `json:"text"`
}
