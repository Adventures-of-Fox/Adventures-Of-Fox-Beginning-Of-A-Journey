package main

import (
	"log"
	"os"
	"server/global"
	"server/options"
)

func main() {
	err,_ := global.CheckOS()
	if err, errMsg := global.CheckError(err, "Starting"); err != nil {
		log.Fatalf(errMsg)
	}
	options.ConfigRun()
	if _, err := os.Stat("server.jar"); err != nil {
		options.Download()
		options.Unzip()
	}
	options.Start()
}
