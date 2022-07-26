package date

import (
	"fmt"
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	s := "20220616"
	tm := ParseDate(s)
	fmt.Printf("%v\n", tm)

	tm = ParseDate("2022-06-16")
	fmt.Printf("%v\n", tm)

	tm = ParseDate("2022/06/16")
	fmt.Printf("%v\n", tm)

	tm = ParseDate("2022-06-16 10:10:10")
	fmt.Printf("%v\n", tm)

	tm = ParseDate("2022年06月16日 10:10:10")
	fmt.Printf("%v\n", tm)
}

func TestParseDatetime(t *testing.T) {
	s := "20220616000001"
	tm := ParseDatetime(s)
	fmt.Printf("%v\n", tm)

	tm = ParseDatetime("2022-06-16 00:02:01")
	fmt.Printf("%v\n", tm)

	tm = ParseDatetime("2022/06/16 00:02:01")
	fmt.Printf("%v\n", tm)

	tm = ParseDatetime("2022-06-16 10:10:10")
	fmt.Printf("%v\n", tm)

	tm = ParseDatetime("2022年06月16日 10时10分10秒")
	fmt.Printf("%v\n", tm)
}

func TestDiffDays(t *testing.T) {
	tm1 := ParseDatetime("2022-06-16 00:02:01")
	tm2 := ParseDatetime("2022-06-17 00:02:00")
	println(DiffDay(tm1, tm2))
	println(DiffDaySec(tm1, tm2))
	println(DiffSec(tm1, tm2))
	println(DiffMin(tm1, tm2))
	println(DiffHour(tm1, tm2))
}

func TestFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(Fmt(now, YmdhmsDash))
	fmt.Println(Fmt(now, YmdDash))
	fmt.Println(Fmt(now, YmdhmsEmpty))
	fmt.Println(Fmt(now, YmdEmpty))
	fmt.Println(Fmt(now, YmdhmsSlash))
	fmt.Println(Fmt(now, YmdSlash))
	fmt.Println(Fmt(now, YmdhmsZh))
	fmt.Println(Fmt(now, YmdZh))
}

func TestFmt1(t *testing.T) {
	i := 86399
	duration := time.Duration(i) * time.Second
	println(int64(duration.Hours()))
	println(int64(duration.Minutes()))
	println(int64(duration.Seconds()))

	h := i / 3600
	m := i % 3600 / 60
	s := i % 3600 % 60
	println(h)
	println(m)
	println(s)
}

func TestAdd(t *testing.T) {
	now := time.Now()
	fmt.Println(FmtDateTime(AddHour(now, 1)))
	fmt.Println(FmtDateTime(SubHour(now, 1)))
}
