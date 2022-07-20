package main

import (
	"fmt"
	"net/url"
)

const myURL = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	//Parsing the URL
	result, _ := url.Parse(myURL) // Parsing the url
	fmt.Println(result.Scheme)    // Prints - https
	fmt.Println(result.Host)      // Prints - www.youtube.com
	fmt.Println(result.Path)      // Prints - /watch
	fmt.Println(result.Query())   // It is map. Prints - map[index:[26] list:[PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa] v:[cl7_ouTMFh0]]

	// Combining parts of the URL into an URL
	urlParts := &url.URL{
		Scheme:   "https",
		Host:     "www.youtube.com",
		Path:     "/watch",
		RawQuery: "index=26&v=cl7_ouTMFh0",
	}

	combinedURL := urlParts.String()
	fmt.Println(combinedURL) // Prints - https://www.youtube.com/watch?index=26&v=cl7_ouTMFh0
}
