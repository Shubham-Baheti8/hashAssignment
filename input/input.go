package input

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func sanitizeInput(input string) string {
	input = strings.ReplaceAll(input, " ", "")
	regex := regexp.MustCompile("[^a-zA-Z0-9]+")
	input = regex.ReplaceAllString(input, "")
	return input
}

func GetSanitizedAlphanumericInput() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter your string: ")
		str, _ := reader.ReadString('\n')
		sanitized := sanitizeInput(str)
		if len(sanitized) == 0 {
			fmt.Println("Invalid input! Please enter alphanumeric characters only.")
		} else {
			return sanitized
		}
	}
}
