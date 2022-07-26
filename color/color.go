package color

import "strconv"

type RGB struct {
	red, green, blue int64
}

func NewRGB(r, g, b int64) RGB {
	return RGB{r, g, b}
}

func t2x(t int64) string {
	result := strconv.FormatInt(t, 16)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}

func (color RGB) Hex() string {
	r := t2x(color.red)
	g := t2x(color.green)
	b := t2x(color.blue)
	return r + g + b
}

func Hex2RGB(hex string) RGB {
	r, _ := strconv.ParseInt(hex[:2], 16, 10)
	g, _ := strconv.ParseInt(hex[2:4], 16, 18)
	b, _ := strconv.ParseInt(hex[4:], 16, 10)
	return RGB{r, g, b}
}
