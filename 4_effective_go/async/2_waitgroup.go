package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	// A slice of sample websites
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}
	var wg sync.WaitGroup

	for i, u := range urls {
		// Increment the wait group counter
		wg.Add(1)
		go func(url string, index int) {
			// Decrement the counter when the go routine completes
			defer wg.Done()
			fmt.Println(index)
			// Call the function check
			checkUrl(url)
		}(u, i)
	}
	// Wait for all the checkWebsite calls to finish
	wg.Wait()
}

//checks and prints a message if a website is up or down
func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down !!!")
		return
	}
	fmt.Println(url, "is up and running.")
}
