package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Text input function for users
func GetUrlInput(prompt string) string {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
