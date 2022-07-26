package tools

import (
	"bytes"
	"encoding/binary"
	"math"
	"normal_pkg/assert"
)

func IntToBytes(n int) ([]byte, error) {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	err := binary.Write(bytebuf, binary.BigEndian, data)
	if err != nil {
		return nil, err
	}
	return bytebuf.Bytes(), nil
}

func BytesToInt(bys []byte) (int64, error) {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	err := binary.Read(bytebuff, binary.BigEndian, &data)
	if err != nil {
		return 0, err
	}
	return data, err
}

func RemainLow(i int64, l int) int64 {
	assert.TrueM(i > 0, "illegal argument")
	assert.TrueM(l > 0, "illegal argument")

	var r int64
	top := int64(math.Pow10(l - 1))
	if i <= top {
		return i
	}
	// 不断除以10，得到低位余数
	for k := int64(10); k <= i; k *= 10 {
		r = i % k
		if k > top {
			break
		}
	}
	return r
}
