package main

import (
	"fmt"
	"net/http"
	"time"
)

var interval int = 3 // seconds
var url string = "http://192.168.0.90:9000/"
var counter int = 0

func sendRequest() {
	fmt.Printf("[%d] Sending request to %s\n", counter, url)

	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("[%d] Error: %s\n", counter, err)
		fmt.Println()
		return
	}

	fmt.Println("Sent")
}

func main() {
	go sendRequest()

	c := time.Tick(time.Duration(interval) * time.Second)

	for _ = range c {
		counter++
		go sendRequest()
	}
}
