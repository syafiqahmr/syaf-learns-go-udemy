package main

import (
	"fmt"
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

	// // an infinite loop
	// for {
	// 	// when receive a link from the channel
	// 	// call checkLink on the same link again
	// 	go checkLink(<-c, c)
	// }

	// wait for the channel to receive some value
	// then assign the value to l
	// more readable then the code block above
	// but achieve the same pupose
	for l := range c {
		// pause current go routine by 2s
		// but this will make the main routine wait for 2s
		// everytime it receives something
		// which is making the main routing slow
		// time.Sleep(time.Second * 2)

		// function literal aka unnamed function
		// php - anon function, python - lambda
		// note the syntax and the parantheses at the end to invoke the function
		// where l is passed into the child routine
		go func(copyLink string) {
			time.Sleep(2 * time.Second)
			// why cannot just accept l
			// this l variable in the child routine
			// will point to the l in the main routine
			// which might be getting changed over time
			// hence NEVER point to the same variable from two diff routines
			// if you need the variable, pass the variable as an argument
			// since go is a pass by value language,
			// the child routine will receive a copy of the value
			checkLink(copyLink, c)
		}(l)
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
	c <- link
}
