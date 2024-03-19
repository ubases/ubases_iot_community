// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package iotutil

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"net"

	"github.com/sony/sonyflake"
	"github.com/speps/go-hashids"

	"cloud_platform/iot_common/iotstrings"
)

var sf *sonyflake.Sonyflake
var upperMachineID uint16

func init() {
	var st sonyflake.Settings
	st.MachineID = getRandomMachineID
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		sf = sonyflake.NewSonyflake(sonyflake.Settings{
			MachineID: lower16BitIP,
		})
		upperMachineID, _ = upper16BitIP()
	}
}

func getRandomMachineID() (uint16, error) {
	for {
		i, err := rand.Int(rand.Reader, big.NewInt(65536))
		if err != nil {
			return 0, err
		}
		mid := i.Uint64()
		if 0 < mid && mid < 65536 {
			return uint16(mid), nil
		}
	}
}

func GetIntID() uint64 {
	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}
	return id
}

func GetUUID(prefix string) string {
	id := GetIntID()
	hd := hashids.NewData()
	h, err := hashids.NewWithData(hd)
	if err != nil {
		panic(err)
	}
	numbers := []int64{int64(id)}
	if upperMachineID != 0 {
		numbers = append(numbers, int64(upperMachineID))
	}
	i, err := h.EncodeInt64(numbers)
	if err != nil {
		panic(err)
	}
	return prefix + iotstrings.Reverse(i)
}

const (
	Alphabet62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	Alphabet36 = "abcdefghijklmnopqrstuvwxyz1234567890"
	Alphabet52 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Number10   = "0123456789"
)

func GetUUID36(prefix string) string {
	id := GetIntID()
	hd := hashids.NewData()
	hd.Alphabet = Alphabet36
	h, err := hashids.NewWithData(hd)
	if err != nil {
		panic(err)
	}
	numbers := []int64{int64(id)}
	if upperMachineID != 0 {
		numbers = append(numbers, int64(upperMachineID))
	}
	i, err := h.EncodeInt64(numbers)
	if err != nil {
		panic(err)
	}
	return prefix + iotstrings.Reverse(i)
}

func randString(letters string, n int) string {
	output := make([]byte, n)
	randomness := make([]byte, n)
	_, err := rand.Read(randomness)
	if err != nil {
		return ""
	}
	l := len(letters)
	for pos := range output {
		random := uint8(randomness[pos])
		randomPos := random % uint8(l)
		output[pos] = letters[randomPos]
	}
	return string(output)
}

func GetSecret(len int) string {
	return randString(Alphabet62, len)
}

func GetAlphaSecret(len int) string {
	return randString(Alphabet52, len)
}

func GetSecretNumber(len int) string {
	return randString(Number10, len)
}

func lower16BitIP() (uint16, error) {
	ip, err := IPv4()
	if err != nil {
		return 0, err
	}
	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}

func upper16BitIP() (uint16, error) {
	ip, err := IPv4()
	if err != nil {
		return 0, err
	}
	return uint16(ip[0])<<8 + uint16(ip[1]), nil
}

func IPv4() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}
		ip := ipnet.IP.To4()
		return ip, nil
	}
	return nil, errors.New("no ip address")
}

//=====================定制方法======================================

// GetProductKeyRandomString 获取产品Key
func GetProductKeyRandomString() string {
	return fmt.Sprintf("PK%s", GetSecret(6))
}

// GetCloudFirmwareKeyRandomString 获取云管固件key
func GetCloudFirmwareKeyRandomString() string {
	key := randString(Alphabet36, 10)
	return fmt.Sprintf("fws%s", key)
}

// GetCustomFirmwareKeyRandomString 获取自定义固件Key
func GetCustomFirmwareKeyRandomString() string {
	key := randString(Alphabet36, 10)
	return fmt.Sprintf("fwc%s", key)
}
