package main

import (
	"fmt"
	"strconv"
)

func IdCardToBirthDay(idCard string) string {
	yr, _ := strconv.Atoi(Substr(idCard, 6, 4))
	month, _ := strconv.Atoi(Substr(idCard, 10, 2))
	day, _ := strconv.Atoi(Substr(idCard, 12, 2))
	return fmt.Sprintf("%d-%02d-%02d", yr, month, day)
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}
