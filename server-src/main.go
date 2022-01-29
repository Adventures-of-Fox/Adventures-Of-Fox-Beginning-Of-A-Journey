package main

import (
	"log"
	"server/global"
	"server/options"
)

func main() {
	err,_ := global.CheckOS()
	if err, errMsg := global.CheckError(err, "Starting"); err != nil {
		log.Fatalf(errMsg)
	}
	options.ConfigRun()
	options.Download()
	options.Unzip()
	options.Install()
}
