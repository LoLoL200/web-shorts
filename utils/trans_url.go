package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func Conclusion(getURL string) {

	apiURL := "https://tinyurl.com/api-create.php?url=" + url.QueryEscape(getURL)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("error with request", err)
		return
	}

	// Working with the client
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error with request")
		return
	}
	defer response.Body.Close()

	// Read the answer
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Read error:", err)
	}

	// Checking
	CheckingStatusCode(response)
	fmt.Printf("Status code: %d\n", response.StatusCode)
	fmt.Println("Original URL:", getURL)
	fmt.Println("Short URL:", string(body))
}
