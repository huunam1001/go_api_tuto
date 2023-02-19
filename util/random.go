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

func RandomString() string {

	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < k; i++ {

		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
