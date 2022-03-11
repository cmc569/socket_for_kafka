package jwt_secret

import (
	"api/util/crypto"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

var jwtSecret []byte

type Claims struct {
	Token string `json:"token"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(usersId int64) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(12000 * time.Hour)
	token, _ := crypto.KeyEncrypt(fmt.Sprint(usersId))
	claims := Claims{
		token,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (usersId int64, err error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			usersIdStr, _ := crypto.KeyDecrypt(claims.Token)
			usersId, _ := strconv.ParseInt(usersIdStr, 10, 64)
			return usersId, nil
		}
	}

	return 0, err
}
