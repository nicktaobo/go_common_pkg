package assert

import "testing"

func TestNotEmptyM(t *testing.T) {
	var a []any
	a = append(a, "   ")
	NotEmptyM(a, "error")
}
