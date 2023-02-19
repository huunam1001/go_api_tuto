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

func RandomInt(min int64, max int64) int64 {

	return min + rand.Int63n(max-min+1)
}

func RandomString(characters int) string {

	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < characters; i++ {

		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwnerName() string {
	return RandomString(6)
}

func RandomBalance() int64 {
	return RandomInt(10, 1000)
}

func RandomCurrency() string {

	currencies := []string{"USD", "EUR", "VND", "JPY"}

	return currencies[rand.Intn(len(currencies))]
}
