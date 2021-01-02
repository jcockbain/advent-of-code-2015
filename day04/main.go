package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("--- Part One ---")
	fmt.Println(Part1("yzbqklnj"))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2("yzbqklnj"))
}

func Part1(input string) int {
	for i := 0; i < 1<<31; i++ {
		inp := input + intToString(i)
		if startsWithZeroes(GetMD5Hash(inp), 5) {
			return i
		}
	}
	return -1
}

func Part2(input string) int {
	for i := 0; i < 1<<31; i++ {
		inp := input + intToString(i)
		if startsWithZeroes(GetMD5Hash(inp), 6) {
			return i
		}
	}
	return -1
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func intToString(n int) string {
	return strconv.Itoa(n)
}

func startsWithZeroes(text string, n int) bool {
	for i := 0; i < n; i++ {
		if string(text[i]) != "0" {
			return false
		}
	}
	return true
}
