package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// Set up the initial client
	client := &http.Client{}

	// Build a GET request
	req, err := http.NewRequest("GET", "https://api.ipify.org?format=text", nil)

	// Perform a request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	// Inside of the get, the http library opens the response Body. We
	// want to make sure that we release the open Body when we're done
	// no matter what happens with out execution (though in this simple
	// example it's not as big of a deal)
	defer resp.Body.Close()

	// Access the body and read it all in, and print it out
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	// Change the Timeout in the client to 5 seconds
	client.Timeout = 5 * time.Second

	// Now attempt to access something that is unavailable:
	//   http://169.254.1.2/example
	req, err = http.NewRequest("GET", "http://169.254.1.2/example", nil)
	resp, err = client.Do(req)
	fmt.Println(err)

}
