package cipher

import (
	"testing"
	"encoding/base64"
)

func TestDes_Encrypt(t *testing.T) {
	origData := []byte("123456")
	key := []byte("woda@Z2y")

	des := new(Des)
	des.Mode = CBCMode
	des.PadType = PKCS5Pad

	cryptData, err := des.Encrypt(origData, key)
	if err != nil {
		t.Fatal(err)
	}

	t.Log( base64.StdEncoding.EncodeToString(cryptData))

	decryptData, err := des.Decrypt(cryptData, key)
	if err != nil {
		t.Fatal(err)
	}

	if string(decryptData) != string(origData) {
		t.Fatal("解密后不相等")
	}

	t.Log(string(decryptData))
}
