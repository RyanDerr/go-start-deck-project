package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		// l is companct syntax to mean wait for a value to be sent from the channel and store in a variable l
		// l is equivilent to <-c and is a blocking call, so the main routine will wait for a value to be sent to the channel
		// Only 5 or the length of the links slice will be go routines will exist at a time
		go func() {
			time.Sleep(getRandomTimeDelay() * time.Second)
			go checkLink(l, c)
		} ()
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	// Sleep The Current Go Routine for A Random Amount of Time Between 1 and 5 Seconds
	// time.Sleep(getRandomTimeDelay() * time.Second)
	c <- link
}

func getRandomTimeDelay() time.Duration {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return time.Duration(random.Intn(5) + 1)
}
