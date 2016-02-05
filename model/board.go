package model

import "math/rand"

const(
	Width			= 10
	Height			= 8

)
var items			= []string{"A","E","I","O","U"}

type Board struct {
	Elements 	[Height][Width]Cell
	enabled		bool
}

func NewBoard() *Board{
	var board = new (Board)
	//board.Elements
	for i:= range board.Elements {
		for j:= range board.Elements[i] {
			board.Elements[i][j].Value = items[rand.Intn(5)]
		}
	}
	return board
}

func (self *Board) SetDefaults(){
	self.enabled	= false
}

type Cell struct {
	Visited 	bool
	Value 		string
}

func (self *Cell) SetDefaults(){
	self.Visited 	= 	false
	self.Value 		= 	"X"
}


