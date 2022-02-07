package main

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/tkanos/gonfig"
)

type configuration struct {
	Interval int
	Url      string
}

var Config *configuration

func InitConfig() {
	fmt.Println(getFileName())

	Config = &configuration{}
	err := gonfig.GetConf(getFileName(), Config)
	if err != nil {
		panic(err)
	}
}

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func getFileName() string {
	filename := "config.json"

	_, dirname, _, _ := runtime.Caller(0)
	filepath := path.Join(filepath.Dir(dirname), filename)

	return filepath
}
