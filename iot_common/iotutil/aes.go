package iotutil

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/forgoer/openssl"
)

//AES加密（CBC模式）
func AES_CBC_EncryptBase64(plainText []byte, key []byte) (str string, err error) {
	defer func() {
		if err0 := recover(); err0 != nil {
			err = errors.New(fmt.Sprintf("%v", err0))
		}
	}()
	iv := []byte("1234567890123456")
	skey := make([]byte, 16)
	copy(skey, key)
	dst, err0 := openssl.AesCBCEncrypt(plainText, skey, iv, openssl.ZEROS_PADDING)
	if err0 == nil {
		str = base64.StdEncoding.EncodeToString(dst)
	}
	return
}

//AES解密（CBC模式）
func AES_CBC_DecryptBase64(cipherText []byte, key []byte) (buf []byte, err error) {
	defer func() {
		if err0 := recover(); err0 != nil {
			err = errors.New(fmt.Sprintf("%v", err0))
		}
	}()
	if buf0, err0 := base64.StdEncoding.DecodeString(string(cipherText)); err0 == nil {
		iv := []byte("1234567890123456")
		skey := make([]byte, 16)
		copy(skey, key)
		dst, err1 := openssl.AesCBCDecrypt(buf0, skey, iv, openssl.ZEROS_PADDING)
		if err1 == nil {
			buf = dst
		} else {
			err = err1
		}
	} else {
		err = err0
	}
	return
}

func AES_CBC_EncryptHex(plainText []byte, key []byte) (str string, err error) {
	defer func() {
		if err0 := recover(); err0 != nil {
			err = errors.New(fmt.Sprintf("%v", err0))
		}
	}()
	iv := []byte("1234567890123456")
	skey := make([]byte, 16)
	copy(skey, key)
	dst, err0 := openssl.AesCBCEncrypt(plainText, skey, iv, openssl.ZEROS_PADDING)
	if err0 == nil {
		str = hex.EncodeToString(dst)
	}
	return
}

//AEC解密（CBC模式）
func AES_CBC_DecryptHex(cipherText []byte, key []byte) (buf []byte, err error) {
	defer func() {
		if err0 := recover(); err0 != nil {
			err = errors.New(fmt.Sprintf("%v", err0))
		}
	}()
	if buf0, err0 := hex.DecodeString(string(cipherText)); err0 == nil {
		iv := []byte("1234567890123456")
		skey := make([]byte, 16)
		copy(skey, key)
		dst, err1 := openssl.AesCBCDecrypt(buf0, skey, iv, openssl.ZEROS_PADDING)
		if err1 == nil {
			buf = dst
		} else {
			err = err1
		}
	} else {
		err = err0
	}
	return
}
