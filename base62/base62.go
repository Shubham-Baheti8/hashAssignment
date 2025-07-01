package base62

import (
	"strings"

	"github.com/Shubham-Baheti8/hashAssignment/config"
)

func EncodeBase62(hash int64) string {
	if hash == 0 {
		return strings.Repeat(config.PaddingChar, config.Base62Len)
	}
	result := ""
	for hash > 0 {
		rem := hash % 62
		hash = hash / 62
		result = string(config.Base62Chars[rem]) + result
	}

	switch {
	case len(result) < config.Base62Len:
		result = strings.Repeat(config.PaddingChar, config.Base62Len-len(result)) + result
	case len(result) > config.Base62Len:
		result = result[len(result)-config.Base62Len:]
	}
	return result
}
