package utils

import (
	"crypto/rand"
	"math/big"
)

func GenerateKey(n int) string {
	result := make([]byte, n)
	base := big.NewInt(int64(len(base62Chars)))

	for i := 0; i < n; i++ {
		randomNum, err := rand.Int(rand.Reader, base)
		if err != nil {
			return ""
		}
		result[i] = base62Chars[randomNum.Int64()]
	}

	return string(result)
}
