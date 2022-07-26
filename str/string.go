package str

import (
	"reflect"
	"strconv"
	"strings"
)

func IsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

func BlurEmail(email string) string {
	p := strings.IndexByte(email, '@')
	if p < 0 {
		return email
	}
	name := email[0:p]
	domain := email[p:len(email)]
	b := strings.Builder{}
	blur := Blur(name, 2, len(name)-1, "*", 4)
	b.WriteString(blur)
	b.WriteString(domain)
	return b.String()
}

func Blur(str string, start int, end int, sep string, num int) string {
	l := len(str)
	if start >= l {
		return str
	}
	b := strings.Builder{}
	prev := str[0:start]
	suf := str[end:l]
	b.WriteString(prev)
	for i := 0; i < num; i++ {
		b.WriteString(sep)
	}
	b.WriteString(suf)
	return b.String()
}

func EndsWith(str string, sep string) bool {
	l := len(str)
	if sep == "" {
		return true
	}
	if l == 0 {
		return false
	}
	i := strings.LastIndex(str, sep)
	return l-i-len(sep) == 0
}

func StartsWith(str string, sep string) bool {
	l := len(str)
	if sep == "" {
		return true
	}
	if l == 0 {
		return false
	}
	i := strings.Index(str, sep)
	return i == 0
}

func String(a any) string {
	var v = reflect.ValueOf(a)
	switch v.Kind() { // 类型的种类Kind，将类型归类，如：int、uint、string、bool、struct、ptr、interface等等
	case reflect.Invalid: // 零值
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10) // 按照十进制int格式化
	case reflect.Uint8, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', 10, 64)
	case reflect.Complex64, reflect.Complex128:
		return strconv.FormatComplex(v.Complex(), 'f', 10, 64)
	case reflect.String:
		// return strconv.Quote(v.String())
		return v.String()
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.Chan, reflect.Map, reflect.Func, reflect.Pointer, reflect.Slice: // 引用类型和函数、指针
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16) // 地址
	default:
		return v.Type().String() + " value" // 默认输出 v的类型Type value
	}
}
