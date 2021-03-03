package model

type Pokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type []Type `json:"type"`
	Base Base   `json:"base"`
}

type Base struct {
	HP        int `json:"HP"`
	Attack    int `json:"Attack"`
	Defense   int `json:"Defense"`
	SPAttack  int `json:"Sp. Attack"`
	SPDefense int `json:"Sp. Defense"`
	Speed     int `json:"Speed"`
}

type Type string

const (
	Bug      Type = "Bug"
	Dark     Type = "Dark"
	Dragon   Type = "Dragon"
	Electric Type = "Electric"
	Fairy    Type = "Fairy"
	Fighting Type = "Fighting"
	Fire     Type = "Fire"
	Flying   Type = "Flying"
	Ghost    Type = "Ghost"
	Grass    Type = "Grass"
	Ground   Type = "Ground"
	Ice      Type = "Ice"
	Normal   Type = "Normal"
	Poison   Type = "Poison"
	Psychic  Type = "Psychic"
	Rock     Type = "Rock"
	Steel    Type = "Steel"
	Water    Type = "Water"
)
