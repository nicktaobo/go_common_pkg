package random

import "testing"

func TestAlpha(t *testing.T) {
	for i := 0; i < 20; i++ {
		r := Alpha(i + 1)
		println(r)
	}
}

func TestAlphaLower(t *testing.T) {
	for i := 0; i < 20; i++ {
		r := AlphaLower(i + 1)
		println(r)
	}
}

func TestAlphaUpper(t *testing.T) {
	for i := 0; i < 20; i++ {
		r := AlphaUpper(i + 1)
		println(r)
	}
}

func TestAlphaNumeric(t *testing.T) {
	for i := 0; i < 20; i++ {
		r := AlphaNumeric(i + 1)
		println(r)
	}
}

func TestNumeric(t *testing.T) {
	for i := 0; i < 20; i++ {
		r := Numeric(i + 1)
		println(r)
	}
}
