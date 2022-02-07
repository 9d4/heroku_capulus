package main

import (
	"fmt"
	"net/http"
	"time"
)

var counter int

func sendRequest() {
	fmt.Printf("[%d] Sending request to %s\n", counter, Config.Url)

	_, err := http.Get(Config.Url)
	if err != nil {
		fmt.Printf("[%d] Error: %s\n", counter, err)
		fmt.Println()
		return
	}

	fmt.Println("Sent")
}

func main() {
	InitConfig()

	go sendRequest()

	c := time.Tick(time.Duration(Config.Interval) * time.Second)

	for range c {
		counter++
		go sendRequest()
	}
}
