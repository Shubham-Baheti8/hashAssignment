package config

const (
	Base62Chars   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	HashInitValue = int64(5381)
	HashMod       = int64(1000000007)
	Base62Len     = 10
	PaddingChar   = "0"
)
