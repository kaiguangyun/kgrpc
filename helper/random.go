package helper

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// random string : A-Z a-z
func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	rs := make([]string, length)

	for start := 0; start < length; start++ {
		t := rand.Intn(2)
		if t == 0 {
			// A-Z
			rs = append(rs, string(rand.Intn(26)+65))
		} else {
			// a-z
			rs = append(rs, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(rs, "")
}

// random number less than n, return int
func RandomInt(n int) int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	return random.Intn(n)
}

// random number between int min-max, return int
func RandomBetween(min, max int) int {
	if min >= max {
		return min
	}

	return RandomInt(max-min+1) + min
}

// random number, return string
func RandomIntString(length int) string {
	result := ""
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < length; i++ {
		result += strconv.Itoa(random.Intn(10))
	}

	return result
}
