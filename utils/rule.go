package utils

import (
	"fmt"
	"net/http"
	"net/url"
)

func CheckingURL(rawURL string) bool {
	parsed, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false

	}
	// Check that the scheme is http or https
	if parsed.Scheme != "http" && parsed.Scheme != "https" {

		return false
	}

	return true
}

func CheckingStatusCode(response *http.Response) {
	// Checking status code
	switch response.StatusCode {
	case 200:
		fmt.Println("✅ All right")
	case 404:
		fmt.Println("❌Not found")
	case 500:
		fmt.Println("❌Server error")
	case 503:
		fmt.Println("❌Server unavailable")
	default:
		fmt.Printf("❌Unknown status %d\n", response.StatusCode)
	}
}
