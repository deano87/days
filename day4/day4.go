package main

import (
	"crypto/md5"
	"fmt"
)

// LoopLimit is used for max iterations = ^int32
const LoopLimit = 2147483647

// Mine receives a key and a hash function and returns the fist number
// that creates a hash that starts with 5 zeros
func Mine(key string, hash func(string) string, term func(string) bool) int32 {
	var i int32
	for i = 0; i < LoopLimit; i++ {
		hashedString := hash(fmt.Sprintf("%s%d", key, i))
		if term(hashedString) {
			return i
		}
	}

	// reached max limit, stop running
	return -1
}

// HashMd5 creates a md5 hash
func HashMd5(s string) string {
	data := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(data))
}

// mining term functions

func termFiveZeros(s string) bool {
	if s[:5] == "00000" {
		return true
	}
	return false
}

func termSixZeros(s string) bool {
	if s[:6] == "000000" {
		return true
	}
	return false
}
