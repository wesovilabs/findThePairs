package game

import (
	"bufio"
	"fmt"
	"os"
	"github.com/wesovi/findThePairs/model"
	"github.com/wesovi/findThePairs/utils"
	"strconv"
)
func startGame() (*model.Player,*model.Board){
	utils.CleanScreen()
	fmt.Print(" Please introduce your full name: ")
	scanner:= bufio.NewScanner(os.Stdin)
	var playerFullName = ""
	for scanner.Scan(){
		playerFullName=scanner.Text()
		break
	}
	return model.NewPlayer(playerFullName),model.NewBoard()
}


func display(player *model.Player,board *model.Board){
	utils.CleanScreen()
	model.DisplayPlayer(player)
	model.DisplayBoard(board)
}

func run(board *model.Board,player *model.Player){
	shouldContinuePlaying:= true
	for shouldContinuePlaying {
		display(player,board)
		shouldContinuePlaying = findPair(player,board)
	}
}

func LaunchGame (){
	var player,board= startGame()
	run(board,player)
}

func askForCell()(int,int){
	return askForRow(),askForColumn()
}

func getCell(board *model.Board) (*model.Cell){
	var row,column = askForCell()
	var cell *model.Cell= checkValue(row,column,board)
	if(cell.Visited){
		fmt.Println("Sorry but this cell has already been visited, please try with another one.")
		return getCell(board)
	}
	return cell
}

func updateCellStatus(cell1 *model.Cell,cell2 *model.Cell) {
	cell1.Visited = true
	cell2.Visited = true
}

func updatePlayerStatus(player *model.Player,points int){
	player.Score+=points
	player.Tries+=1
}

func checkIfChosenCellsAreValid(cell1,cell2 *model.Cell)(bool,int){
	if(cell1.Value == cell2.Value){
		return true,10;
	}else{
		return false,-2
	}
}

func findPair(player *model.Player,board *model.Board)(bool) {
	fmt.Println("Select 1st cell:")
	fmt.Println("----------------")
	var cell1 = getCell(board)
	fmt.Println("Select 2nd cell:")
	fmt.Println("----------------")
	var cell2 = getCell(board)
	for (cell1==cell2){
		fmt.Println("Please don't try cheating me... and don't selec twice the same cell.")
		cell2 = getCell(board)
	}
	var _,points =checkIfChosenCellsAreValid(cell1,cell2)
	updatePlayerStatus(player,points)
	updateCellStatus(cell1,cell2)
	return true
}


func askForPoint(name string, minValue,maxValue int) (int,bool){
	fmt.Print(name+"["+strconv.Itoa(minValue)+"-"+strconv.Itoa(maxValue)+"]: ")
	return utils.CheckOption(minValue,maxValue)
}

func askForRow() int{
	var value,valid=askForPoint("Row",1,model.Width)
	if(!valid){
		fmt.Println("Please introduce a valid value for the row ["+strconv.Itoa(1)+"-"+strconv.Itoa(model.Width)+"].")
		return askForRow()

	}
	return value
}

func askForColumn() int{
	var value,valid=askForPoint("Column",1,model.Height)
	if(!valid){
		fmt.Println("Please introduce a valid value for the column ["+strconv.Itoa(1)+"-"+strconv.Itoa(model.Height)+"].")
		return askForColumn()
	}
	return value
}

func checkValue(row int,column int, board *model.Board) *model.Cell{
	return &board.Elements[row-1][column-1]
}
