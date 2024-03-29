package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.stackoverflow.com/",
		"https://www.golang.org/",
		"https://www.amazon.com/",
	}

	c := make(chan string)

	for _, link := range links {
		// asynchronous function call
		go checkLink(link, c)
	}
	
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<- c)
	// }

	// infinite loop old way
	// for {
	// 	go checkLink(<- c, c)
	// }

	// infinite loop new way
	for l := range c {
		go func() {
			time.Sleep(2 * time.Second)
			checkLink(l, c)
		}()
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if (err) != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep it's up"
}