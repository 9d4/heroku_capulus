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
