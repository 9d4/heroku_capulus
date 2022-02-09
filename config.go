package main

import (
	"fmt"
	"os"
	"path"

	"github.com/tkanos/gonfig"
)

type configuration struct {
	Interval string
	Urls     []string
}

var Config *configuration

func InitConfig() {
	fmt.Println("[config path]", getFileName())

	Config = &configuration{}
	err := gonfig.GetConf(getFileName(), Config)
	if err != nil {
		panic(err)
	}

	// check is urls are empty, then panic
	checkUrls()
}

func getFileName() string {
	filename := "config.json"

	dirname, err := os.Getwd()
	if err != nil {
		panic("error getting pwd")
	}

	filepath := path.Join(dirname, filename)

	return filepath
}

func checkUrls() {
	if len(Config.Urls) == 0 {
		panic("Urls are empty")
	}
}