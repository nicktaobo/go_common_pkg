package json

import (
	"bytes"
	"encoding/json"
)

func ToJson(t any) string {
	r, err := json.Marshal(t)
	if err != nil {
		panic(err)
		return ""
	}
	return string(r)
}

func ToJsonF(t any) string {
	r, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		panic(err)
		return ""
	}
	return string(r)
}

func Parse(str string, t any) any {
	err := json.Unmarshal([]byte(str), t)
	if err != nil {
		panic(err)
	}
	return t
}

// ParseUseNum 解析json，如果有数字，使用 Number 处理而不是 float64，防止解析为科学计数法
func ParseUseNum(str string, t any) {
	d := json.NewDecoder(bytes.NewReader([]byte(str)))
	d.UseNumber()
	err := d.Decode(t)
	if err != nil {
		panic(err)
	}
}
