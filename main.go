package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/traperwaze/heroku_capulus/config"
	"github.com/traperwaze/heroku_capulus/ntp"
)

var counter int
var ntpTime time.Time

func main() {
	ntpTime := ntp.GetTime()
	fmt.Println(ntpTime)


	interval, err := time.ParseDuration(config.Config.Interval)
	if err != nil {
		panic("Error parsing interval. Change it in config.json. e.g 1m, 1h10m, 10s, etc.")
	}

	c := time.Tick(interval)

	go func() {
		// this line makes requests without waiting first interval
		sendRequests()

		for range c {
			counter++
			sendRequests()
		}
	}()

	select {}
}

func sendRequest(url string, count int) {
	fmt.Printf("[%d] Sending request to %s\n", count, url)

	res, err := http.Head(url)
	if err != nil {
		fmt.Printf("[%d] Error: %s\n", count, err)
		return
	}

	fmt.Printf("[%d] Sent to %s. status:%d\n", count, url, res.StatusCode)
}

func sendRequests() {
	for _, url := range config.Config.Urls {
		go sendRequest(url, counter)
	}
}
