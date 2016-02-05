package menu

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"github.com/wesovi/findThePairs/model"
	"github.com/wesovi/findThePairs/utils"

)


func cleanScreen(){
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}



func startGame() (*model.Player,*model.Board){
	cleanScreen()
	fmt.Println(" Please introduce your full name.")
	scanner:= bufio.NewScanner(os.Stdin)
	var playerFullName = ""
	for scanner.Scan(){
		playerFullName=scanner.Text()
		break
	}
	return model.NewPlayer(playerFullName),model.NewBoard()
}

func run(board *model.Board,player *model.Player) (){
	var shouldContinuePlaying bool = true
	for shouldContinuePlaying {
	cleanScreen()
	fmt.Println(" "+player.FullName+" has "+strconv.Itoa(player.Score)+" pointss.")
	fmt.Println()
	fmt.Println("   1  2  3  4  5  6  7  8  9  10")
	fmt.Println("  ------------------------------")
	for i:= range board.Elements{
		fmt.Print(strconv.Itoa(i+1) +" ")
		for j:=range board.Elements[i]{

			if(!board.Elements[i][j].Visited){
				//fmt.Print(" "+strconv.Itoa(i)+","+strconv.Itoa(j)+" ")
				fmt.Print( " ? ")
			}else{
				fmt.Print( " "+board.Elements[i][j].Value+" ")
			}
		}
		fmt.Println()
	}
	findPair(player,board)
	}

}

func LaunchGame (){
	var player,board= startGame()
	run(board,player)
}

func findPair(player *model.Player,board *model.Board) {
	var row,column,row2,column2 int
	row = askForRow()
	column = askForColumn()
	var cell1 *model.Cell=checkValue(row,column,board)

	row2 = askForRow()
	column2 = askForColumn()
	var cell2 *model.Cell=checkValue(row2,column2,board)

	if(cell1.Value == cell2.Value){
		player.Score+=10
	}

	cell1.Visited = true
	cell2.Visited = true
	fmt.Println(cell1.Value+ " and "+cell2.Value)
}



func askForRow() int{
	fmt.Print("Row: ")
	var value,valid =utils.CheckOption(1,10)
	if(!valid){
		askForRow()
	}
	return value
}

func askForColumn() int{
	fmt.Print("Column: ")
	var value,valid =utils.CheckOption(1,7)
	if(!valid){
		askForColumn()
	}
	return value
}

func checkValue(row int,column int, board *model.Board) *model.Cell{
	return &board.Elements[column-1][row-1]
}
