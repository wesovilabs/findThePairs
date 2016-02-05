package model

import "math/rand"

const(
	width			= 10
	height			= 8

)
var items			= []string{"A","E","I","O","U"}

type Board struct {
	Elements 	[height][width]Cell
	enabled		bool
}

func NewBoard() *Board{
	var board = new (Board)
	//board.Elements
	for i:= range board.Elements {
		for j:= range board.Elements[i] {
			board.Elements[i][j].Value = items[rand.Intn(5)]
			//board.Elements[i][j].Visited = true
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


