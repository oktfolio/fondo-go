package auth

import (
	"fmt"
	"fondo-go/config"
	"fondo-go/db"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims ...
type Claims struct {
	User *db.User
	jwt.StandardClaims
}

// SignToken 签发token
func SignToken(claims Claims, duration time.Duration, key string) (stoken string, err error) {
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(duration).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    config.Conf.Env.Token.Issuer,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key)) // 必须转换为[]byte, 否则就是无效key
}

// ParseToken 解析token
func ParseToken(stoken, key string) *Claims {
	if stoken == "" || key == "" {
		return nil
	}
	var custom Claims
	token, err := jwt.ParseWithClaims(stoken, &custom, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected sign method %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		log.Println("parse token:", err)
		return nil
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims
	}
	return nil
}
