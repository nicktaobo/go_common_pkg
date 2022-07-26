package prob

import (
	"fmt"
	"math/rand"
	"normal_pkg/assert"
)

// Percent 给定义一个100以内的数字，匹配百分比概率，命中返回 true，否则返回 false
func Percent(r int) bool {
	assert.TrueM(r >= 0 && r <= 100, "invalid parameter")
	if r == 0 {
		return false
	}
	if r == 100 {
		return true
	}
	n := rand.Intn(100)
	return n <= r
}

// Half 概率命中与否的概率相同，都为50%
func Half() bool {
	return rand.Intn(2) == 1
}

// Select 给定一组数据，按照概率选择数据中匹配的下标，数据的值作为概率值，返回匹配到的数组的下标
func Select(is []int) int {
	var num int
	for _, v := range is {
		if v < 0 {
			panic(fmt.Sprintf("invalid parameter, not negative value is required: %d", v))
		}
		num += v
	}
	if num == 0 {
		panic("invalid parameter, total value is zero")
	}
	r := rand.Intn(num)
	rate := 0
	for i, v := range is {
		rate += v
		if r < rate {
			return i
		}
	}
	panic("unknown error")
}
