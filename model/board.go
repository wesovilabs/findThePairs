package model

import(
	"math/rand"
	"fmt"
	"strconv"
)

const(
	Width			= 6
	Height			= 6
)

type Item struct {
	Value  			string
	Perms			int
}

func NewItem(value string) *Item{
	item := new (Item)
	item.Value = value
	item.Perms = 0
	return item
}

type Board struct {
	Elements 	[Height][Width]Cell
	Items 		[]*Item
}


func NewBoard() *Board{
	random := rand.New(rand.NewSource(99))
	var board = new (Board)
	board.Items = []*Item{NewItem("X"),NewItem("O")}
	for i:= range board.Elements {
		for j:= range board.Elements[i] {
			var randomPosition = random.Intn(len(board.Items))
			var item *Item = board.Items[randomPosition]
			board.Elements[i][j].Value = item.Value
			item.Perms+=1
			board.Elements[i][j].Row = i
			board.Elements[i][j].Column = j
		}
	}
	return board
}

type Cell struct {
	Visited 	bool
	Value 		string
	Row 		int
	Column 		int
}

func DisplayPendingItems(board *Board){
	for i:= range board.Items{
		fmt.Println((board.Items[i]).Value+": "+strconv.Itoa(board.Items[i].Perms))
	}
}


func DisplayBoard(board *Board){
	fmt.Println()
	DisplayPendingItems(board)
	fmt.Println()
	fmt.Println("   1  2  3  4  5  6")
	fmt.Println("  -----------------")
	for i:= range board.Elements{
		fmt.Print(strconv.Itoa(i+1) +" ")
		for j:=range board.Elements[i]{
			displayCell(board.Elements[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func displayCell(cell Cell){
	if(!cell.Visited){
		fmt.Print( " - ")
	}else{
		fmt.Print( " "+cell.Value+" ")
	}
}

