package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func makeRequest(url string) {
	var err interface{}

	_, err = http.Get(url)

	defer wg.Done()

	if err != nil {
		panic("Can't make request")
	}
}

func printInColor(r, g, b int, msg string) {
	fmt.Printf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r, g, b, msg)
}

func main() {

	timer := time.NewTicker(time.Millisecond)
	done := make(chan bool, 1)
	var timeItTook int = 0 // in ms
	url := flag.String("url", "", "Enter the url")

	flag.Parse()

	wg.Add(1)
	go makeRequest(*url)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-timer.C:
				timeItTook++
				fmt.Print(".") // loader
			}
		}
	}()

	wg.Wait()
	done <- true
	timer.Stop()
	fmt.Println(" \n ")
	r, g, b := 50, 230, 120
	msg := "Time it took = " + strconv.Itoa(timeItTook) + "ms"
	printInColor(r, g, b, msg)
	fmt.Println(" \n ")
}
