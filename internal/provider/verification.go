package provider

import (
	"math/rand"
	"time"
)

func GenerateVerificationCode() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	max := 999999
	min := 100000
	return r.Intn(max-min) + min
}
