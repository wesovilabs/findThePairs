package utils

import(
	"runtime"
)

const (
	noRealOS		= "windows"
)

func CheckOS()(string,bool){
	var myOs = runtime.GOOS
	return myOs, (myOs!=noRealOS)
}