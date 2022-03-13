package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/9d4/heroku_capulus/config"
)

var (
	interval time.Duration
)

func init() {
	ival, err := time.ParseDuration(config.Config.Interval)
	if err != nil {
		log.Fatal("Error parsing interval. Change it in config.json. e.g 1m, 1h10m, 10s, etc.")
	}

	interval = ival
}

// should be called with go
func req(url string, start chan bool) {
	fmt.Printf("Creating routine: %s\n", url)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	<-start
	fmt.Printf("Start routine: %s\n", url)

	for {
		res, err := http.Head(url)
		if err != nil {
			continue
		}
		
		fmt.Printf("Sent to %s. status:%d\n", url, res.StatusCode)

		<-ticker.C
	}

}

func reqAll() {
	var rs []*chan bool

	for _, url := range config.Config.Urls {
		start := make(chan bool)
		rs = append(rs, &start)

		go req(url, start)
	}

	// start all routines once they're created
	for _, start := range rs {
		*start <- true
	}
}
