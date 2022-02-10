package main

import (
	"fmt"
	"net/http"
	"time"
)

var counter int

func main() {
	InitConfig()

	interval, err := time.ParseDuration(Config.Interval)
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
	for _, url := range Config.Urls {
		go sendRequest(url, counter)
	}
}
