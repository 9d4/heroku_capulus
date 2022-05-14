package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
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

func main() {
	reqAll()
	select {}
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

	afterStart := now.Hour() >= int(startHour) && now.Minute() >= int(startMinute)
	beforeStop := now.Hour() <= int(stopHour) && now.Minute() < int(stopMinute)

	return !(afterStart && beforeStop)
}

func yieldErrorParseTime(err error) {
	log.Fatal("error parsing time", err)
}

func isAlwaysOnMode() bool {
	return config.Config.AlwaysOn
}

// should be called with go
func req(url string, start chan bool, done *sync.WaitGroup) {
	fmt.Printf("Creating routine: %s\n", url)

	tick := time.Tick(interval)
	count := 0

	done.Done()
	<-start
	fmt.Printf("Start routine: %s\n", url)

	for {
		if !isAlwaysOnMode() && isPaused() {
			<- tick
			continue
		}

		res, err := http.Head(url)
		if err != nil {
			fmt.Printf("[%d] %d -- %s\n", count, err, url)
			<- tick
			continue
		}

		count++
		fmt.Printf("[%d] %d -- %s\n", count, res.StatusCode, url)

		<-tick
	}
}

func reqAll() {
	var wg sync.WaitGroup
	var rs []*chan bool

	for _, url := range config.Config.Urls {
		wg.Add(1)
		start := make(chan bool)
		rs = append(rs, &start)

		u := url

		go req(u, start, &wg)
	}

	wg.Wait()

	// start all routines once they're created
	for _, start := range rs {
		*start <- true
		close(*start)
	}
}
