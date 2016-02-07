package model
import (
	"strconv"
	"fmt"
)

type Player struct {
	FullName 	string
	Score		int
	Tries		int
}

func (self *Player) SetDefaults(){
	self.Score		=	0
	self.Tries		= 	0
}

func NewPlayer(fullName string) *Player{
	player := new (Player)
	player.FullName = fullName
	return player
}

func makePlayer(fullName string) Player{
	return Player{fullName,0,0}
}

func DisplayPlayer(player *Player){
	fmt.Println(" "+player.FullName+" has "+strconv.Itoa(player.Score)+" points and has already spent "+strconv.Itoa(player.Tries)+" tries of "+strconv.Itoa(Width * Height /2))
}