package config

import (
	"log"
	"os"
	"path"

	"github.com/tkanos/gonfig"
)

type configuration struct {
	Interval  string
	Urls      []string
	NtpServer string
	Timezone  string
	StartAt   string
	StopAt    string
}

var Config *configuration
var filename string = "config.json"

func init() {
	InitConfig()
}

func InitConfig() {
	Config = &configuration{}

	err := gonfig.GetConf(getFileName(), Config)
	if err != nil {
		log.Fatal(err)
	}

	if urlsEmpty() {
		log.Fatal("[config error] urls is empty")
	}
}

func getFileName() string {
	filepath := path.Join(wd(), filename)

	return filepath
}

func urlsEmpty() bool {
	return len(Config.Urls) == 0
}

func wd() string {
	p, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return p
}
