package hash

import "github.com/Shubham-Baheti8/hashAssignment/config"

func GenerateHash(s string) int64 {
	hash := config.HashInitValue
	for i := 0; i < len(s); i++ {
		hash = ((hash * 33) + int64(s[i])) % config.HashMod
	}
	return hash
}
