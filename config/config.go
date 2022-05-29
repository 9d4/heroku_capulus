package config

import (
	"bufio"
	"log"
	"os"
	"path"

	"github.com/tkanos/gonfig"
)

type configuration struct {
	Interval string
	Urls     []string
	Timezone string
	AlwaysOn bool
	StartAt  string
	StopAt   string
}

var Config *configuration
var filename string = "config.json"

func init() {
	if Config != nil {
		return
	}

	Config = new(configuration)

	err := gonfig.GetConf(getFileName(), Config)
	if err != nil {
		log.Fatal(err)
	}

	if urlsEmpty() {
		log.Fatal("[config error] urls is empty")
	}
}

func initConfigToml() {
	path := path.Join(wd(), "config.toml")
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		println(err)
	}

	buf := bufio.NewScanner(file)

	for buf.Scan() {
		print(buf.Bytes())
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
