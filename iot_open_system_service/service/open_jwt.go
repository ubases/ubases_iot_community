package service

import (
	"cloud_platform/iot_open_system_service/config"
	"errors"

	"github.com/golang-jwt/jwt"

	"time"
)

// type JWTConfig struct {
// 	SigningKey      string
// 	AccessTokenTTL  time.Duration
// 	RefreshTokenTTL time.Duration
// }

type OpenUserInfo struct {
	UserID      int64  `json:"userId"`
	Nickname    string `json:"nickName"`
	Avatar      string `json:"avatar"`
	TenantId    string `json:"tenantId"`
	AccountType int32  `json:"accountType"`
}

type OpenClaims struct {
	jwt.StandardClaims
	OpenUserInfo
}

func (c *OpenClaims) GenerateToken() (tokenStr string, err error) {
	c.IssuedAt = time.Now().Unix()
	c.ExpiresAt = time.Now().Add(time.Second * time.Duration(config.Global.CloudJwt.AccessTokenTTL)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(config.Global.CloudJwt.SigningKey))
}

func (c *OpenClaims) GenerateRefreshToken() (tokenStr string, err error) {
	c.IssuedAt = time.Now().Unix()
	c.ExpiresAt = time.Now().Add(time.Second * time.Duration(config.Global.CloudJwt.RefreshTokenTTL)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(config.Global.CloudJwt.SigningKey))
}

func (c *OpenClaims) ParseToken(tokenStr string) (claims *OpenClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Global.CloudJwt.SigningKey), nil
	})
	if err != nil {
		return &OpenClaims{}, err
	}
	if err := token.Claims.Valid(); err != nil {
		return &OpenClaims{}, err
	}
	if claim, ok := token.Claims.(*OpenClaims); ok {
		return claim, nil
	}
	return &OpenClaims{}, errors.New("Parse errors")
}

func (c *OpenClaims) RefreshToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Global.CloudJwt.SigningKey), nil
	})
	if err != nil {
		return "", err
	}
	if err := token.Claims.Valid(); err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*OpenClaims)
	if !ok {
		return "", errors.New("Parse errors")
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(config.Global.CloudJwt.RefreshTokenTTL)).Unix()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Global.CloudJwt.SigningKey))
}

func (c *OpenClaims) VerifyToken(tokenStr string) bool {
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
