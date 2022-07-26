package collection

import (
	"fmt"
	"normal_pkg/test"
	"reflect"
	"testing"
)

var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestRetain(t *testing.T) {
	ret := NewSlice(slice).Retain(func(a int) bool {
		return a > 5
	})
	expect := []int{6, 7, 8, 9}
	test.Log(t, reflect.DeepEqual(expect, ret), "expect %v, actual %v", expect, ret)

	ret1 := NewSlice([]any{"a", 1, 3.14}).Retain(func(a any) bool {
		switch a.(type) {
		case int:
			return a.(int) > 1
		default:
			return true
		}
	})
	expect1 := []any{"a", 3.14}
	test.Log(t, reflect.DeepEqual(expect1, ret1), "expect %v, actual %v", expect1, ret1)
}

func TestJoin(t *testing.T) {
	s := NewSlice(slice).Join(",")
	test.Log(t, s == "1,2,3,4,5,6,7,8,9", "%v join result should be %s", slice, s)
	join := NewSlice([]string{"a", "b"}).Join(".")
	test.Log(t, join == "a.b", "a join b should be %s", join)
	join = NewSlice([]any{"a", 1, "3.14"}).Join(",")
	test.Log(t, join == "a,1,3.14", "%s join %d join %.2f should be %s", "a", 1, 3.14, join)
}

func TestSort(t *testing.T) {
	var slice = []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	s := NewSlice(slice)
	sr := s.Sort(func(i, j *int) bool {
		if *i < *j {
			return true
		}
		return false
	})
	fmt.Printf("%v\n", slice)
	sr.Reverse()
	fmt.Printf("%v\n", slice)
}

type user struct {
	name  string
	age   int
	score float32
}

func TestSortObj(t *testing.T) {
	users := []user{
		{name: "huzhou", age: 18, score: 99.5},
		{name: "huzhou", age: 16, score: 100},
		{name: "zhangsan", age: 17, score: 99.5},
		{name: "abbc", age: 17, score: 99.5},
	}

	s := NewSlice(users)
	sr := s.Sort(func(a, b *user) bool {
		if a.score != b.score {
			return a.score > b.score
		}
		if a.age != b.age {
			return a.age < b.age
		}
		if a.name != b.name {
			return a.name < b.name
		}
		return false
	})
	fmt.Printf("%v\n", s)
	sr.Reverse()
	fmt.Printf("%v\n", s)
}
