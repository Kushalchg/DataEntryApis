package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type MyClaims struct {
	email string
	cat   string
	role  string
	id    uint
	jwt.StandardClaims
}

func GenerateAccessToken(uemail, urole string, uid uint) (string, error) {
	// load key from env file
	key := []byte(os.Getenv("KEY"))
	accessClaims := MyClaims{
		email: uemail,
		cat:   "access",
		role:  urole,
		id:    uid,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	// pointer to the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	// generate token string using key
	AT, err := token.SignedString(key)

	return AT, err
}
func GenerateRefreshToken(uemail, urole string, uid uint) (string, error) {
	// load key from env file
	key := []byte(os.Getenv("KEY"))
	accessClaims := MyClaims{
		email: uemail,
		cat:   "refresh",
		role:  urole,
		id:    uid,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 50).Unix(),
		},
	}
	// pointer to the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	// generate token string using key
	AT, err := token.SignedString(key)

	return AT, err
}

func ParseToken(token string) (*MyClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("KEY")), nil
	})

	return parsedAccessToken.Claims.(*MyClaims), err
}
