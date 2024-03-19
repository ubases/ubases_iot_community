package iotstruct

import (
	"time"

	jsoniter "github.com/json-iterator/go"
)

type UcUserOperate struct {
	Id         int64     `json:"id"`         // 主键ID
	TenantId   string    `json:"tenantId"`   // 租户id
	AppKey     string    `json:"appKey"`     // appKey
	UserId     int64     `json:"userId"`     // 用户ID(t_uc_user.id)
	Account    string    `json:"account"`    // 登录账号
	RequestUri string    `json:"requestUri"` // 请求地址
	Ip         string    `json:"ip"`         // ip
	OptTime    time.Time `json:"optTime"`    // 操作时间
}

func (u UcUserOperate) MarshalBinary() ([]byte, error) {
	return jsoniter.Marshal(u)
}

func (u *UcUserOperate) UnmarshalBinary(data []byte) error {
	return jsoniter.Unmarshal(data, u)
}

type OpenUserLogin struct {
	UserId    int64  `json:"userId"`    // 用户ID
	TenantId  string `json:"tenantId"`  // 租户id
	LoginTime int64  `json:"time"`      // 登录时间
	ExpiresAt int64  `json:"expiresAt"` //过期时间
	Ip        string `json:"ip"`        // 登录IP
	Addr      string `json:"addr"`      //登录地址
	Token     string `json:"token"`     //访问token
}

func (u OpenUserLogin) MarshalBinary() ([]byte, error) {
	return jsoniter.Marshal(u)
}

func (u *OpenUserLogin) UnmarshalBinary(data []byte) error {
	return jsoniter.Unmarshal(data, u)
}
