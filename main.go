package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/traperwaze/heroku_capulus/config"
)

var counter int

func main() {
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
		if !isPaused() {
			go sendRequest(url, counter)
		}
	}
}

func isPaused() bool {
	tz, err := time.LoadLocation(config.Config.Timezone)
	if err != nil {
		log.Fatal("Error loading timezone:", err)
	}

	now := time.Now().In(tz)

	startAt := strings.Split(strings.Trim(config.Config.StartAt, " "), ":")
	stopAt := strings.Split(strings.Trim(config.Config.StopAt, " "), ":")

	startHour, err := strconv.ParseInt(startAt[0], 10, 64)
	if err != nil {
		yieldErrorParseTime(err)
	}

	startMinute, err := strconv.ParseInt(startAt[1], 10, 64)
	if err != nil {
		yieldErrorParseTime(err)
	}

	stopHour, err := strconv.ParseInt(stopAt[0], 10, 64)
	if err != nil {
		yieldErrorParseTime(err)
	}

	stopMinute, err := strconv.ParseInt(stopAt[1], 10, 64)
	if err != nil {
		yieldErrorParseTime(err)
	}

	afterStart := now.Hour() >= int(startHour) || now.Minute() >= int(startMinute)
	beforeStop := now.Hour() <= int(stopHour) || now.Minute() < int(stopMinute)

	return !(afterStart && beforeStop)
}

func yieldErrorParseTime(err error) {
	log.Fatal("error parsing time", err)
}
