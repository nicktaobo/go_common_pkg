package collection

import (
	"normal_pkg/str"
	"reflect"
	"sort"
	"strings"
)

// Slice is a generic slice type
type Slice[T any] []T

func NewSlice[T any](s []T) Slice[T] {
	if reflect.TypeOf(s).Kind() != reflect.Slice {
		panic("require slice")
	}
	return s
}

// Retain is a method that retain the elements that matched the condition the function param defined,
// and not matched elements will be removed.
// This method will return a new slice and the original slice will not be changed
func (s Slice[T]) Retain(cond func(a T) bool) []T {
	var ret []T
	for _, a := range s {
		if cond(a) { // 符合条件
			ret = append(ret, a)
		}
	}
	return ret
}

// Filter is a method that opposite to the Retain method, it will filter the elements that not matched the condition
// the function param defined, and not matched elements will be retained.
// This method will return a new slice and the original slice will not be changed
func (s Slice[T]) Filter(cond func(a T) bool) []T {
	var ret []T
	for _, a := range s {
		if !cond(a) { // 不符合条件
			ret = append(ret, a)
		}
	}
	return ret
}

// Join is a method that join all the elements by the string the sep param defined
func (s Slice[T]) Join(sep string) string {
	var ret []string
	for _, a := range s {
		ret = append(ret, str.String(a))
	}
	return strings.Join(ret, sep)
}

// SortableSlice is a struct to define a sortable slice, it implements sort.Interface.
type SortableSlice[T any] struct {
	slice Slice[T]
	less  func(x, y *T) bool // 比较的方法，参数为T的指针，直接更改原始slice的顺序
}

func (s SortableSlice[T]) Len() int {
	return len(s.slice)
}

func (s SortableSlice[T]) Less(i, j int) bool {
	return s.less(&s.slice[i], &s.slice[j])
}

func (s SortableSlice[T]) Swap(i, j int) {
	s.slice[i], s.slice[j] = s.slice[j], s.slice[i]
}

// Sort is a method to sort the original slice by the given argument less, which is a method to define how to compare
// the elements in it, this less method receive two *T type so Sort method will change the original slice.
// Sort will return a SortableSlice instance, so you could invoke its other api such as SortableSlice.Reverse etc.
func (s Slice[T]) Sort(less func(x, y *T) bool) SortableSlice[T] {
	v := &SortableSlice[T]{s, less}
	sort.Sort(v)
	return *v
}

// Reverse is a method to reverse the sequence that the Slice.Sort method sorted, so you should has invoked the Slice.Sort
// first before to call Reverse.
func (s SortableSlice[T]) Reverse() Slice[T] {
	sort.Sort(sort.Reverse(s))
	return s.slice
}
