package utils

import(
	"runtime"
	"os/exec"
	"os"
)

const (
	noRealOS		= "windows"
)

func CheckOS()(string,bool){
	var myOs = runtime.GOOS
	return myOs, (myOs!=noRealOS)
}

func CleanScreen(){
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}