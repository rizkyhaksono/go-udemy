package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://www.google.com/")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// bs := make([] byte, 99999)
	// resp.Body.Read(bs)

	lw := logWriter{}

	io.Copy(lw, resp.Body)
	// fmt.Println(io.Copy(os.Stdout, resp.Body))

	// fmt.Println(string(bs))
}

func (logWriter) Write (bs[] byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))

	return len(bs), nil
}