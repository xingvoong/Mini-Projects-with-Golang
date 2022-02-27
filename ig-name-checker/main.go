package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	usernames := []string{
		"barackobama",
		"MichelleObama",
		"elonmusk",
		"12fdfvscc.",
		"xinvoongacxczxccxxc  12e3e",
	}

	// No concurrency
	fmt.Println("No concurrency")
	startNoConcurrency := time.Now()

	for _, username := range usernames {
		checkUsernameNoConcurrency(username)
	}
	fmt.Println(time.Since(startNoConcurrency))
	fmt.Println()

	// WaitGroup
	fmt.Println("Wait Group")
	start := time.Now()

	for _, username := range usernames {
		go checkUsername(username)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
	fmt.Println()

	// Channel
	fmt.Println("Channel")
	startChannel := time.Now()
	c := make(chan string, len(usernames))

	for _, username := range usernames {
		go checkUsernameChannel(username, c)
		wg.Add(1)
	}

	wg.Wait()
	close(c)

	for status := range c {
		fmt.Println(status)
	}
	fmt.Println(time.Since(startChannel))
}

func checkUsernameNoConcurrency(username string) {

	igUser := "https://instagram.com/" + username + "?__a=1"

	if res, err := http.Get(igUser); err != nil {
		fmt.Printf("IG username %s does not exist \n", username)
	} else {

		if res.StatusCode == 404 {
			fmt.Printf("IG username %s does not exist \n", username)
		} else {
			fmt.Printf("[%d] IG username %s exists \n", res.StatusCode, username)
		}

	}

}

func checkUsername(username string) {
	defer wg.Done()

	igUser := "https://instagram.com/" + username + "?__a=1"

	if res, err := http.Get(igUser); err != nil {
		fmt.Printf("IG username %s does not exist \n", username)
	} else {

		if res.StatusCode == 404 {
			fmt.Printf("IG username %s does not exist \n", username)
		} else {
			fmt.Printf("[%d] IG username %s exists \n", res.StatusCode, username)
		}

	}

}

func checkUsernameChannel(username string, c chan string) {
	defer wg.Done()

	igUser := "https://instagram.com/" + username + "?__a=1"

	if res, err := http.Get(igUser); err != nil {
		c <- fmt.Sprintf("IG username %s does not exist \n", username)
	} else {
		if res.StatusCode == 404 {
			c <- fmt.Sprintf("IG username %s does not exist \n", username)
		} else {
			c <- fmt.Sprintf("[%d] IG username %s exists", res.StatusCode, username)
		}

	}
}
