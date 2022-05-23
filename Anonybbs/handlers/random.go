package handlers

import (
	"math/rand"
	"time"
)

func GenerateRandomString() string {
	var r = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0")
	uniq := make([]rune, 12)
	for i := range uniq {
		rand.Seed(time.Now().UTC().UnixNano() + int64(i<<12))
		uniq[i] = r[rand.Intn(len(r))]
	}
	return string(uniq)
}
