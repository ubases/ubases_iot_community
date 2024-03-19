package service

import (
	"cloud_platform/iot_system_service/config"
	"errors"

	"github.com/golang-jwt/jwt"

	"time"
)

type JWTConfig struct {
	SigningKey      string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

type CloudUserInfo struct {
	UserID   int64  `json:"userId"`
	Nickname string `json:"nickName"`
	Avatar   string `json:"avatar"`
	DeptId   int64  `json:"deptId"`  //云平台用户才有
	RoleIds  string `json:"roleIds"` //云平台用户才有，多个用逗号分隔
	PostIds  string `json:"postIds"` //云平台用户才有，多个用逗号分隔
}

type CloudClaims struct {
	jwt.StandardClaims
	CloudUserInfo
}

func (c *CloudClaims) GenerateToken() (tokenStr string, err error) {
	c.IssuedAt = time.Now().Unix()
	c.ExpiresAt = time.Now().Add(time.Second * time.Duration(config.Global.CloudJwt.AccessTokenTTL)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(config.Global.CloudJwt.SigningKey))
}

func (c *CloudClaims) GenerateRefreshToken() (tokenStr string, err error) {
	c.IssuedAt = time.Now().Unix()
	c.ExpiresAt = time.Now().Add(time.Second * time.Duration(config.Global.CloudJwt.RefreshTokenTTL)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(config.Global.CloudJwt.SigningKey))
}

func (c *CloudClaims) ParseToken(tokenStr string) (claims *CloudClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Global.CloudJwt.SigningKey), nil
	})
	if err != nil {
		return &CloudClaims{}, err
	}
	if err := token.Claims.Valid(); err != nil {
		return &CloudClaims{}, err
	}
	if claim, ok := token.Claims.(*CloudClaims); ok {
		return claim, nil
	}
	return &CloudClaims{}, errors.New("Parse errors")
}

func (c *CloudClaims) RefreshToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Global.CloudJwt.SigningKey), nil
	})
	if err != nil {
		return "", err
	}
	if err := token.Claims.Valid(); err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*CloudClaims)
	if !ok {
		return "", errors.New("Parse errors")
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(config.Global.CloudJwt.RefreshTokenTTL)).Unix()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Global.CloudJwt.SigningKey))
}

func (c *CloudClaims) VerifyToken(tokenStr string) bool {
	token, err := jwt.ParseWithClaims(tokenStr, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Global.CloudJwt.SigningKey), nil
	})
	if err != nil {
		return false
	}
	if err := token.Claims.Valid(); err != nil {
		return false
	}
	return true
}
