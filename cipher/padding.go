package cipher

import (
	"bytes"
)

func PKCS5Padding(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)% blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, padText...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
