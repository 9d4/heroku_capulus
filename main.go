package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var counter int

func sendRequest() {
	url := os.Getenv("URL")

	fmt.Printf("[%d] Sending request to %s\n", counter, url)

	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("[%d] Error: %s\n", counter, err)
		fmt.Println()
		return
	}

	fmt.Printf("[%d] Sent\n", counter)
}

func main() {
	InitEnv()

	go sendRequest()

	interval, err := time.ParseDuration(os.Getenv("INTERVAL"))
	if err != nil {
		panic("Error parsing interval. Change it in .env")
	}

	c := time.Tick(interval)

	for range c {
		counter++
		go sendRequest()
	}
}
