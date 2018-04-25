package cipher

import (
	"crypto/cipher"
	cryptoDes "crypto/des"
	"errors"
)

//TODO others padding type and mode

const (
	PKCS5Pad = 0

	CBCMode = 0
)

type Des struct {
	PadType		int
	Mode		int
}


func (des *Des) Encrypt(origData, key []byte) ([]byte, error){
	block, err := cryptoDes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	switch des.PadType {
	case PKCS5Pad:
		origData = PKCS5Padding(origData, block.BlockSize())
	default:
		return nil, errors.New("PadType not support")
	}

	cryptData := make([]byte, len(origData))

	switch des.Mode {
	case CBCMode:
		blockMode := cipher.NewCBCEncrypter(block, key)
		blockMode.CryptBlocks(cryptData, origData)
	default:
		return nil, errors.New("Mode not support")
	}

	return cryptData, nil
}

func (des *Des) Decrypt(cryptData, key []byte) ([]byte, error) {
	block, err := cryptoDes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	origData := make([]byte, len(cryptData))

	switch des.Mode {
	case CBCMode:
		blockMode := cipher.NewCBCDecrypter(block, key)
		blockMode.CryptBlocks(origData, cryptData)
	default:
		return nil, errors.New("Mode not support")
	}

	switch des.PadType {
	case PKCS5Pad:
		origData = PKCS5UnPadding(origData)
	default:
		return nil, errors.New("PadType not support")
	}

	return origData, nil
}
