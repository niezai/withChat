package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JwtSecret = []byte("2023-12-05 yellow i love you so")

type Claims struct {
	UserUuid string `json:"user_uuid"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 签发token
func GenerateToken(userId string, userName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		UserUuid: userId,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "FanOne-Mail",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret)
	return token, err
}

// 验证用户 获取用户信息
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
