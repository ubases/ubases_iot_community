/**
 * @Author: hogan
 * @Date: 2022/3/23 19:44
 */
package iotutil

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

//rand seed如何初始化
//rand.Seed(time.Now().UnixNano())

// 生成随机float
func RandFloat() string {
	return strconv.FormatFloat(rand.Float64(), 'f', 6, 64)
}

// gen uuid
func Uuid() string {
	id, _ := uuid.NewUUID()
	return strings.ReplaceAll(id.String(), "-", "")
}

// to md5
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// to sha1
func Sha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Base64Encoding base64 加密
func Base64Encoding(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64Decoding base64 解密
func Base64Decoding(s string) string {
	decodeStr, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return string(decodeStr)
}

// GetRandomString 随机生成字符串
func GetRandomString(l int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GetRandomPureString 随机生成纯字符串
func GetRandomPureString(l int) string {
	str := "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GetRandomNumber 随机生成数字字符串
func GetRandomNumber(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
