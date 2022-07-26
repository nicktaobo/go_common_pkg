package random

import (
	"math/rand"
	"strings"
	"time"
)

var (
	lowerChars    = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	upperChars    = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	chars         = append(lowerChars, upperChars...)
	numbers       = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	alphanumerics = append(append(lowerChars, numbers...), upperChars...)
)

func AlphaLower(le int) string {
	return randomString(le, lowerChars)
}

func AlphaUpper(le int) string {
	return randomString(le, upperChars)
}

func Alpha(le int) string {
	return randomString(le, chars)
}

func Numeric(le int) string {
	return randomString(le, numbers)
}

func AlphaNumeric(le int) string {
	return randomString(le, alphanumerics)
}

func randomString(le int, data []byte) string {
	if le <= 0 {
		panic("length must be > 0")
	}
	rand.Seed(time.Now().UnixNano())
	b := strings.Builder{}
	for i := 0; i < le; i++ {
		idx := rand.Intn(len(data))
		b.WriteByte(data[idx])
	}
	return b.String()
}
