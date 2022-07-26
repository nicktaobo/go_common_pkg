package assert

import (
	"fmt"
	"strings"
)

func True(b bool) {
	if !b {
		panic(fmt.Sprintf("assert failed, expected [%t] but found [%t]", true, b))
	}
}

func TrueM(b bool, msg string) {
	if !b {
		panic(msg)
	}
}

func False(b bool) {
	if b {
		panic(fmt.Sprintf("assert failed, expected [%t] but found [%t]", false, b))
	}
}

func FalseM(b bool, msg string) {
	if b {
		panic(msg)
	}
}

func Nil(t any) {
	if t != nil {
		panic(fmt.Sprintf("assert failed, expected [nil] but found [not nil]: %v", t))
	}
}

func NilM(t any, msg string) {
	if t != nil {
		panic(msg)
	}
}

func NotNil(t any) {
	if t == nil {
		panic(fmt.Sprintf("assert failed, expected [not nil] but found [nil]"))
	}
}

func NotNilM(t any, msg string) {
	if t == nil {
		panic(msg)
	}
}

func NotBlank(s string) {
	s = strings.TrimSpace(s)
	if s == "" {
		panic(fmt.Sprintf("assert failed, expected [not empty] but found [\"\"]"))
	}
}

func NotBlankM(s string, msg string) {
	s = strings.TrimSpace(s)
	if s == "" {
		panic(msg)
	}
}

func Equals(t1 any, t2 any) {
	if t1 != t2 {
		panic(fmt.Sprintf("assert failed, expected t1 equals t2 but not"))
	}
}

func EqualsM(t1 any, t2 any, msg string) {
	if t1 != t2 {
		panic(msg)
	}
}

func NotEmpty(d []any) {
	if d == nil || len(d) == 0 {
		panic(fmt.Sprintf("assert failed, expected none nil and none empty collection"))
	}
}

func NotEmptyM(d []any, msg string) {
	if d == nil || len(d) == 0 {
		panic(msg)
	}
}

func ErrorM(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}
