package cryptox

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

func JWTEncode(key string, claims jwt.MapClaims) (string, error) {
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenObj.SignedString([]byte(key))
}

func JWTDecode(key string, tokenString string) (jwt.MapClaims, error) {
	// 解析 token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// 验证加密方式是否一致
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(key), nil
	})
	// 解析失败
	if err != nil {
		return nil, err
	}
	// 校验 token 是否有效
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	// 断言 claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot convert to MapClaims")
	}
	return claims, nil
}
