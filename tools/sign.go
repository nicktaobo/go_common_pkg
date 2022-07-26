package tools

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func PKCS5Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	if padding == 0 {
		return plaintext
	}
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

// PKCS5UnPadding 去除填充数据
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AESEncryptToECB(s string, key string) string {
	data := []byte(s)
	keyb := []byte(key)

	block, _ := aes.NewCipher(keyb)
	// AES分组长度为128位，所以blockSize=16，单位字节
	bs := block.BlockSize()
	data = PKCS5Padding(data, bs)

	if len(data)%bs != 0 {
		panic("need multiple block size")
	}

	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return fmt.Sprintf("%x", out)
}

func AESDecryptToECB(s string, key string) string {
	data, _ := hex.DecodeString(s)
	keyb := []byte(key)

	block, _ := aes.NewCipher(keyb)
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		panic("need multiple block size")
	}

	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return string(out)
}

func AESEncryptToCBC(s string, key string) string {
	src := []byte(s)
	keyb := []byte(key)

	block, _ := aes.NewCipher(keyb)
	// AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	src = PKCS5Padding(src, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, keyb[:blockSize]) // 初始向量的长度必须等于块block的长度16字节
	crypted := make([]byte, len(src))
	blockMode.CryptBlocks(crypted, src)
	// hex编码
	return fmt.Sprintf("%x", crypted)
}

func AESDecryptToCBC(s string, key string) string {
	// Hex解码
	src, _ := hex.DecodeString(s)
	keyb := []byte(key)

	block, _ := aes.NewCipher(keyb)
	// AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyb[:blockSize]) // 初始向量的长度必须等于块block的长度16字节
	origData := make([]byte, len(src))
	blockMode.CryptBlocks(origData, src)
	origData = PKCS5UnPadding(origData)
	return string(origData)
}

func intToBytes(i int) []byte {
	x := int32(i)
	buff := bytes.NewBuffer([]byte{})
	binary.Write(buff, binary.BigEndian, x)
	return buff.Bytes()
}
