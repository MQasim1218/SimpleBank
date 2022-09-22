package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
}

func randNum(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randStr(slen int) string {
	var stb strings.Builder
	alphs := "abcdefghijklmnopqrstuvwxyz"
	l := len(alphs)

	for i := 0; i < slen; i++ {
		stb.WriteByte(alphs[rand.Intn(l)])
	}

	return stb.String()
}

func RandomCurrency() string {
	curr := []string{"USD", "GBP", "EUR", "TKY"}
	return curr[rand.Intn(len(curr))]
}

func RandonMoney() int64 {
	return randNum(1000, 10000)
}

func RandName() string {
	return randStr(rand.Intn(12))
}
