package main

import (
	"fmt"
	"github.com/wesovi/findThePairs/menu"
	"github.com/wesovi/findThePairs/utils"
)

func main(){
	var _, valid = utils.CheckOS()
	if(!valid){
		fmt.Println("See you when using a real os, XDD")
	}else{
		menu.LaunchGame()
	}

}
