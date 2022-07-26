package conv

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// StringToInt covert string to int
func StringToInt(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

// StringToInt64 covert string to int64
func StringToInt64(str string) int64 {
	n, _ := strconv.ParseInt(str, 10, 64)
	return n
}

// IntToString covert int to string
func IntToString(src int) string {
	return strconv.Itoa(src)
}

// Int64ToString covert int64 to string
func Int64ToString(src int64) string {
	return strconv.FormatInt(src, 10)
}

func BigIntArrayToString(src []*big.Int) string {
	var temp = make([]string, len(src))
	for k, v := range src {
		temp[k] = fmt.Sprintf("%d", v.Int64())
	}
	var result = strings.Join(temp, ",")
	return result
}

func StringToFloat64(amount string) float64 {
	float, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0
	}
	return float
}

//func ParseFloat64(s string, bitSize int) (float64, error) {
//	if bitSize == 32 {
//		f, err := strconv.tof32(s)
//		return float64(f), err
//	}
//	return atof64(s)
//}

func Int64To16ScaleString(src int64) string {
	return strconv.FormatInt(src, 16)
}
func String16ToInt64(src string) int64 {
	id, err := strconv.ParseUint(src, 16, 32)
	if err != nil {
		panic(err)
	}
	return int64(id)
}

func SplitStringToInt(s string, sec string) []int64 {
	ss := strings.Split(s, sec)
	var is []int64
	for _, i := range ss {
		it, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			panic(err)
		}
		is = append(is, it)
	}
	return is
}

func IntsToString(is []int64) []string {
	if is == nil || len(is) == 0 {
		return nil
	}
	var ss = make([]string, len(is))
	for _, i := range is {
		ss = append(ss, strconv.FormatInt(i, 10))
	}
	return ss
}

func StrsToInt(ss []string) []int64 {
	if ss == nil || len(ss) == 0 {
		return nil
	}
	var is = make([]int64, len(ss))
	for _, s := range ss {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil
		}
		is = append(is, i)
	}
	return is
}
