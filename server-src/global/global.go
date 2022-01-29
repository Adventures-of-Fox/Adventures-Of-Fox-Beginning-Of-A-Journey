package global

import (
	"errors"
	"log"
	"runtime"

	"github.com/fatih/color"
)

func CheckError(err error, name string) (error, string) {
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		red := color.New(color.FgRed, color.Bold)
		return err, red.Sprintf("[%v] %v", name, color.RedString(err.Error()))
	}
	return nil, ""
}

func Logs(msg string, other ...interface{}) {
	info := color.New(color.FgCyan)
	log.Print(info.Sprintf(msg, other...))
}

func Done(msg string, other ...interface{}) {
	info := color.New(color.FgGreen, color.Bold)
	log.Print(info.Sprintf(msg, other...))
}

func CheckOS() (error,bool) {
	switch runtime.GOOS {
	case "windows":
		return nil,false
	case "darwin":
	case "linux":
		return nil,true
	default:
		err := errors.New("That operation system is not supported.")
		return err,false
	}
	return nil,false
}

