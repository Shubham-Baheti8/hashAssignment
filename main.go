package main

import (
	"fmt"

	"github.com/Shubham-Baheti8/hashAssignment/base62"
	"github.com/Shubham-Baheti8/hashAssignment/hash"
	"github.com/Shubham-Baheti8/hashAssignment/input"
)

func main() {
	userInput := input.GetSanitizedAlphanumericInput()
	hashValue := hash.GenerateHash(userInput)
	base62String := base62.EncodeBase62(hashValue)
	fmt.Printf("Final 10-char Base62 hash: %s\n", base62String)
}
