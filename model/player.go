package model

type Player struct {
	FullName 	string
	Score		int
}

func (self *Player) SetDefaults(){
	self.Score		=	0
}

func NewPlayer(fullName string) *Player{
	player := new (Player)
	player.FullName = fullName
	return player
}

func makePlayer(fullName string) Player{
	return Player{fullName,0}
}