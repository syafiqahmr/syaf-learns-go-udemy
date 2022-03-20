package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// makes channel of type string
	c := make(chan string)

	for _, link := range links {
		// this method is slow
		// we are waiting for the response from the first link
		// before proceeding to the next one
		// checkLink(link)

		// make use of concurrency
		// note you can only use go keyword in front of function call
		// however, the main routine does not care that the child routine
		// is not done, hence when you just do this,
		// you wont see anything getting printed out
		// as the main routine will exit before the child routines are done
		go checkLink(link, c)
	}

	// this for loop will prevent the main routine from exiting
	// before all the child routines are done
	for i := 0; i < len(links); i++ {
		// this is a blocking call
		// and will wait for it to receive something
		// aka main routine will sleep
		// before receiving something
		fmt.Println(<-c)
	}
}

// as the channel created in main is only accessible there
// you need to pass the channel as argument
// else check link won't have access to the channel
// note the syntax, and how you must also specity type
func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down i think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep its up"
}

// Go Routine executes our compiled programme line by line
// A blocking call is any function that takes time
// go checkLink(link)
// by adding the go keyword before the function call,
// a brand new Go Routine is created and runs that function
// inside checkLink, when it hits the blocking call (http.Get)
// the programme goes to sleep and the main routine continues
// in essence, the go keyword creates a new thread go routine

// Go Scheduler:
// cpu
// ^
// scheduler
// ^
// go routine, go routine
// scheduler runs one routine until it finishes or make a blocking call
// despite there being multiple routines
// HOWEVER
// when you have multiple cpu cores,
// schedule runs on thread on each "logical" core
// by default go tries to use one core, unless specified

// concurrency vs parallelism
// concurrency - we can have multiple threads executing code,
// if one thread blocks, another one is picked up and worked on
// aka your programme can have multiple threads,
// but only one is worked on at any given time
// parallelism - multiple threads executed at the exact same time
// requires multiple cpus

// when running a programme, a default main routine is created
// child routines is created by the 'go' keyword
// however the main routine always have priority

// channel is the only way you can communicate between go routines
// channel of type string can only communicate with strings and so on

// sending data with cahnnels:
// 1. channel <- 5 : send the value '5' into this channel
// 2. myNumber <- channel : wait for a value to be sent into the channel,
// when we get one assign the value to 'myNumber'
// 3. fmt.Println(<-channel) : wait for a value to be sent into the channel,
// when we get one log it out immediately
// * if you send something to the channel but no routine is receiving it,
// the programme will not exit as it is waiting for someone to receive it
