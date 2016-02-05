package utils

import (
	"bufio"
	"os"
	"fmt"
)


func CheckOption(min int, max int) (int,bool){
	var intValue int
	var value string
	scanner:= bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		value=scanner.Text()
		break
	}
	if _, err := fmt.Sscan(value, &intValue); err != nil {
		return -1,false
	}
	return intValue,intValue<=max&&intValue>=min
}