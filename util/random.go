package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generate random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var result strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		letter := alphabet[rand.Intn(k)]
		result.WriteByte(letter)
	}

	return result.String()
}

func RandomOwner() string {
	return RandomString(8)
}

func RandomMoney() int64 {
	return RandomInt(0, 10000000)
}

func RandomCurrency() string {
	currencies := []string{"IDR", "USD", "JPY", "UER", "SAR"}
	k := len(currencies)

	return currencies[rand.Intn(k)]
}
