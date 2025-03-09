package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// Vrátí sub claim JWT tokenu.
//
// @param tokenStr
// @return uint64
// @return error
func GetUserIdFromToken(tokenStr string) (uint, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwt.MapClaims{})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}
	userId := claims["sub"].(float64)
	return uint(userId), nil
}

// Generate and sign access token.
//
//	@param id
//	@return string
//	@return error
func GenerateAccessToken(id uint) (string, error) {
	// token lifespan
	// payload
	now := time.Now()
	claims := jwt.MapClaims{}
	claims["sub"] = id
	claims["type"] = "user"
	claims["exp"] = now.Add(time.Second * time.Duration(GetSingleton().Config.AccessTokenLifespan)).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// create and sign token
	return token.SignedString([]byte(GetSingleton().Config.AccessTokenSecret))
}
