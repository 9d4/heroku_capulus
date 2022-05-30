package config

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/pelletier/go-toml/v2"
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

	err := initConfigToml(Config)
	if err != nil {
		err = initConfigJson(Config)
	}

	if err != nil {
		log.Fatal("[config error] no valid config found")
	}


	if urlsEmpty() {
		log.Fatal("[config error] urls is empty")
	}
}

func initConfigToml(cfg *configuration) error {
	path := path.Join(wd(), "config.toml")
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	doc,err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = toml.Unmarshal(doc, cfg)
	if err != nil {
		return err
	}

	return nil
}

func initConfigJson(cfg *configuration) error {
	err := gonfig.GetConf(getFileName(), cfg)
	return err
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
