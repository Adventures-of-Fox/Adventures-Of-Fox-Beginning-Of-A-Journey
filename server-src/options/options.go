package options

import (
	//"regexp"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"server/global"
	"server/modules"
	"strings"
)

func S(text string, args ...interface{}) string {
	return fmt.Sprintf(text, args...)
}

func R(text string, arg string) string {
	return strings.ReplaceAll(text, arg, "")
}

func Download() {
	modules.Download("https://meta.fabricmc.net/v2/versions/loader/1.18.1/0.12.12/0.10.2/server/jar", "Downloading server jar", "server.jar")
	modules.Download("https://media.forgecdn.net/files/3616/627/Mineshafts++Monsters+-+1.10.9.zip", "Downloading mods", ".tpm/server-mods.zip")
}

func Unzip() {
	modules.Unziping(".tpm/server-mods.zip", "mods")
}

func Start() {
	global.Logs("Starting server")
	commad :=  S(
		`%v -Xmx%vM -Xms%vM -jar server.jar nogui`, 
		ConfigRun().Java, 
		ConfigRun().RamMaxMB,
		ConfigRun().RamMinMB,
	)
	if _, unix := global.CheckOS(); unix {
		global.Logs(commad)
		modules.RunCmdLive(ConfigRun().Shell, "-c", commad)
	} else {
		global.Logs(commad)
		modules.RunCmdLive(commad)
	}
	global.Done("Server has been stopped")
}

type ConfigJSON struct {
	Java     string
	RamMinMB int
	RamMaxMB int
	Shell    string
}

func ConfigRun() ConfigJSON {
	var config ConfigJSON
	_, err := os.Stat("config.json")
	if err != nil {
		global.Logs("Creating config file")
		if _, unix := global.CheckOS(); unix {
			cmd := modules.RunCmdGet(os.Getenv("SHELL"), "-c", `which java`)
			config = ConfigJSON{
				Java:     R(cmd, "\n"),
				RamMinMB: 4096,
				RamMaxMB: 4096,
				Shell:    os.Getenv("SHELL"),
			}
		} else {
			cmd := modules.RunCmdGet("where", "java")
			config = ConfigJSON{
				Java:     R(cmd, "\n"),
				RamMinMB: 4096,
				RamMaxMB: 4096,
			}
		}
		jsonMarshall, err := json.Marshal(config)
		if err, errMsg := global.CheckError(err, "Creating config failed"); err != nil {
			log.Fatalf(errMsg)
		}
		err = ioutil.WriteFile("config.json", jsonMarshall, 4077)
		if err, errMsg := global.CheckError(err, "Creating config failed"); err != nil {
			log.Fatalf(errMsg)
		}
		global.Done("Created config file")
		return config
	}
	file, err := ioutil.ReadFile("config.json")
	if err, errMsg := global.CheckError(err, "Reading config failed"); err != nil {
		log.Fatalf(errMsg)
	}
	err = json.Unmarshal(file, &config)
	if err, errMsg := global.CheckError(err, "Reading config failed"); err != nil {
		log.Fatalf(errMsg)
	}
	return config
}
