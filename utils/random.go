package utils

import (
	"math/rand"
	"strings"
	"time"
)

const characters = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().Unix())
}

// Function to create a andom integer between a given range
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Function to create a random string given its length
func RandomString(n int) string {
	var sb strings.Builder
	k := len(characters)

	for i := 0; i < n; i++ {
		randChar := characters[rand.Intn(k)]
		sb.WriteByte(randChar)
	}

	return sb.String()
}

// Generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// Generates random amount
func RandomBalance() int64 {
	return RandomInt(0, 1000)
}

// Generates random currency from given three currencies
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
