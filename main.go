package main

import (
	"fmt"
	"web-shorts/utils"
)

func main() {
	for {
		fmt.Println("🔗 URL Shortener")
		fmt.Println("Enter URL to shorten (or 'exit' to exit):")
		url := utils.GetUrlInput(">")
		if url == "exit" {
			fmt.Println("👋 Goodbye")
			return
		}
		fmt.Println("📡 I am sending a request to TinyURL API...")
		if !utils.CheckingURL(url) {
			fmt.Println("💀Invalid URL")
			return
		}
		utils.Conclusion(url)

	}

}
