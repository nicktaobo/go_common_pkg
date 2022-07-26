package json

import (
	"fmt"
	"testing"
	"time"
)

func TestToJson(t *testing.T) {
	var mp = make(map[string]any)
	var smp = make(map[string]any)
	mp["a"] = 1
	mp["b"] = 2

	smp["s1"] = 11
	smp["s2"] = 22
	mp["c"] = smp

	println(ToJson(mp))
}

type User struct {
	Name      string    // `json:"name"`
	Age       int       // `json:"age"`
	BirthDate time.Time // `json:"birthDate"`
}

func TestToJson1(t *testing.T) {
	var user = User{
		Name:      "张三",
		Age:       20,
		BirthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}
	println(ToJsonF(user))
}

func TestToJsonf(t *testing.T) {
	var user = User{
		Name:      "张三",
		Age:       20,
		BirthDate: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
	}
	println(ToJson(user))
}

func TestParse(t *testing.T) {
	s := "{\"Name\":\"张三\",\"Age\":20,\"BirthDate\":\"2000-01-01T00:00:00+08:00\"}"
	u := Parse(s, &User{})
	fmt.Printf("user: %v\n", u)
}

func TestParse1(t *testing.T) {
	s := "{\n  \"Name\": \"张三\",\n  \"Age\": 20,\n  \"BirthDate\": \"2000-01-01T00:00:00+08:00\"\n}"
	u := Parse(s, &User{})
	fmt.Printf("user: %v\n", u)
}

func TestParse2(t *testing.T) {
	s := "{\"Name\":\"张三\",\"Age\":20,\"BirthDate\":\"2000-01-01T00:00:00+08:00\"}"
	var mp = make(map[string]any)
	Parse(s, &mp)
	fmt.Printf("user: %v\n", mp)
	fmt.Printf("name: %v", mp["Name"])
	fmt.Printf("age: %v", mp["Age"])
	fmt.Printf("date: %v", mp["BirthDate"])
}
