package main

import (
	"fmt"
	"net/http"
	"time"
)

var counter int

func sendRequest(url string) {
	fmt.Printf("[%d] Sending request to %s\n", counter, url)

	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("[%d] Error: %s\n", counter, err)
		return
	}

	fmt.Printf("[%d] Sent to %s\n", counter, url)
}

func sendRequests() {
	for _, url := range Config.Urls {
		go sendRequest(url)
	}
}

func main() {
	InitConfig()

	// this line makes requests without waiting first interval
	sendRequests()

	interval, err := time.ParseDuration(Config.Interval)
	if err != nil {
		panic("Error parsing interval. Change it in config.json. e.g 1m, 1h10m, 10s, etc.")
	}

	c := time.Tick(interval)

	for range c {
		counter++
		sendRequests()
	}
}
