package test

import "testing"

func TestError(t *testing.T) {
	Failed(t, "%s should equals %s", "abc", "123")
	Failed(t, "should equals")
}

func TestPass(t *testing.T) {
	Pass(t, "%s should equals %s", "abc", "abc")
	Pass(t, "should equals")
}
