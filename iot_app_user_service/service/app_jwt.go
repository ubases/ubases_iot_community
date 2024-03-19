package service

import (
	"cloud_platform/iot_app_user_service/config"
	"errors"

	"github.com/golang-jwt/jwt"

	"time"
)

type AppUserInfo struct {
	UserID         int64  `json:"userId"`
	Nickname       string `json:"nickName"`
	Avatar         string `json:"avatar"`
	Account        string `json:"account"`
	TenantId       string `json:"tenantId"`
	AppKey         string `json:"appKey"`
	RegionServerId int64  `json:"regionServerId"`
	//HomeIds  string `json:"homeIds"`
}

type AppClaims struct {
	jwt.StandardClaims
	AppUserInfo
}

func (c *AppClaims) GenerateToken() (tokenStr string, err error) {
	c.IssuedAt = time.Now().Unix()
	c.ExpiresAt = time.Now().Add(time.Second * time.Duration(config.Global.AppJwt.AccessTokenTTL)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(config.Global.AppJwt.SigningKey))
}

func (c *AppClaims) GenerateRefreshToken() (tokenStr string, err error) {
	c.IssuedAt = time.Now().Unix()
	c.ExpiresAt = time.Now().Add(time.Second * time.Duration(config.Global.AppJwt.RefreshTokenTTL)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(config.Global.AppJwt.SigningKey))
}

func (c *AppClaims) ParseToken(tokenStr string) (claims *AppClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Global.AppJwt.SigningKey), nil
	})
	if err != nil {
		return &AppClaims{}, err
	}
	if err := token.Claims.Valid(); err != nil {
		return &AppClaims{}, err
	}
	if claim, ok := token.Claims.(*AppClaims); ok {
		return claim, nil
	}
	return &AppClaims{}, errors.New("Parse errors")
}

func (c *AppClaims) RefreshToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Global.AppJwt.SigningKey), nil
	})
	if err != nil {
		return "", err
	}
	if err := token.Claims.Valid(); err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*AppClaims)
	if !ok {
		return "", errors.New("Parse errors")
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(config.Global.AppJwt.RefreshTokenTTL)).Unix()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Global.AppJwt.SigningKey))
}

func (c *AppClaims) VerifyToken(tokenStr string) bool {
	token, err := jwt.ParseWithClaims(tokenStr, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Global.AppJwt.SigningKey), nil
	})
	if err != nil {
		return false
	}
	if err := token.Claims.Valid(); err != nil {
		return false
	}
	return true
}
