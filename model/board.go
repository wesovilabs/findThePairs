package model

import(
	"math/rand"
	"fmt"
	"strconv"
	"time"
)

const(
	Width			= 6
	Height			= 6

)

var PossibleElements		= []string{"A","B","C"}

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
	Items 		map[string]*Item
}


func initializeBoard(board *Board){
	board.Items = make(map[string]*Item)
	for _,element:= range PossibleElements {
		board.Items[element] = NewItem(element)
	}
}

func NewBoard() *Board{
	rand.Seed( time.Now().UTC().UnixNano())
	var board = new (Board)
	initializeBoard(board)
	for i:= range board.Elements {
		for j:= range board.Elements[i] {
			var randomPosition = rand.Intn(len(board.Items))
			var item *Item = board.Items[PossibleElements[randomPosition]]
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
	fmt.Println("Hidden letters:")
	for _,option:= range PossibleElements{
		fmt.Println(" - "+(board.Items[option]).Value+": "+strconv.Itoa(board.Items[option].Perms))
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

