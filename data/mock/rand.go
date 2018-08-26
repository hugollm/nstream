package mock

import "math/rand"

const LETTERS = "abcdefghijklmnopqrstuvwxyz"

func RandString(length int) string {
	bts := make([]byte, length)
	for i := range bts {
		bts[i] = LETTERS[rand.Intn(len(LETTERS))]
	}
	return string(bts)
}
