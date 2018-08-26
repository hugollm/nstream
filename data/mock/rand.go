package mock

import (
	"math/rand"
	"time"
)

const LETTERS = "abcdefghijklmnopqrstuvwxyz"

var seeded bool

func RandString(length int) string {
	seedOnce()
	bts := make([]byte, length)
	for i := range bts {
		bts[i] = LETTERS[rand.Intn(len(LETTERS))]
	}
	return string(bts)
}

func seedOnce() {
	if !seeded {
		rand.Seed(time.Now().UTC().UnixNano())
		seeded = true
	}
}
