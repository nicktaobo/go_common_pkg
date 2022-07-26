package tools

import (
	"encoding/base64"
	json2 "encoding/json"
	"fmt"
	"testing"
)

func TestAES(t *testing.T) {
	// mcnKXZcY/6o5Bv7czf5sfg==
	// test数据
	key := "1585903878056926"
	s := "{\"sysKey\":\"test_key\",\"sysValue\":\"test_value1\",\"name\":\"测试\",\"description\":\"测试描述\"}"
	ret := AESEncryptToCBC(s, key)
	fmt.Println("AESEncryptToCBC: ", ret)
	ret = AESDecryptToCBC(ret, key)
	fmt.Println("AESDecryptToCBC: ", ret)

	ret = AESEncryptToECB(s, key)
	fmt.Println("AESEncryptToECB: ", ret)
	ret = AESDecryptToECB(ret, key)
	fmt.Println("AESEncryptToECB: ", ret)

	var params map[string]string
	json2.Unmarshal([]byte(s), &params)
	fmt.Printf("%v\n", params)
}

func TestBase64(t *testing.T) {
	json := "{\"uuid\": 1654759783138}"
	s := base64.StdEncoding.EncodeToString([]byte(json))
	println(s)
	s = base64.URLEncoding.EncodeToString([]byte(json))
	println(s)
}

func TestAes1(t *testing.T) {
	s := "hello + world = \"hello, world\"!"
	key := "1111111111111111"
	// 833cd28843dacd0b86d4e7f40d1531c51032031b5516d2cff86df4d45af323e4
	d := AESEncryptToECB(s, key)
	// 9f3cb0fa12cd8b980990a9d23a8115114faeafdc2bcc14d5b09b6f276837ace8
	println(d)
}
