package tools

import (
	"fmt"
	"testing"
)

func TestMergeMap(t *testing.T) {
	m1 := make(map[string]string)
	m1["a"] = "a"
	m1["b"] = "b"
	m2 := make(map[string]string)
	m1["c"] = "c"
	m1["d"] = "d"
	m1["a"] = "e"

	MergeMap(m1, m2)
	fmt.Println(len(m1))
	fmt.Printf("merge: %v\n", m1)
}

type Test struct {
	Id int
}

func fun() (Test, error) {
	return Test{123}, nil
}

func TestMust(t *testing.T) {
	t1 := Must(fun())
	fmt.Println(t1.Id)
}
