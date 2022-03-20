package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	fileName := os.Args[1]
	f, e := os.Open(fileName)

	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, f)

}

func (lw logWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}
