package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	// 1. http request to http
	// 2. print the response

	// resp here is *Response type
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// make a byte slice with 99999 empty elements
	// pass this bs and data into the Read function
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	// does the same thing as the previous part
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// now the logWritter is implementing the Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	// need to check documentation to udnerstand
	// what you are suppose to return
	return len(bs), nil
}

// other notes:
// 1. if a struct has a field with type interface, it means that
// any type works so long as it is of the specified interface
// 2. interfaces can also put together multiple interfaces eg.
// type ReadCloser interface{
// 	Reader
// 	Closer
// }
