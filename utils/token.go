package utils

import (
	"dbs2/models"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Generate and sign access token.
//
//	@param id
//	@return string
//	@return error
func GenerateAccessToken(user *models.User) (string, error) {
	// token lifespan
	// payload
	now := time.Now()
	claims := jwt.MapClaims{}
	claims["sub"] = user.ID
	claims["type"] = user.Role
	claims["exp"] = now.Add(time.Second * time.Duration(GetSingleton().Config.AccessTokenLifespan)).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// create and sign token
	return token.SignedString([]byte(GetSingleton().Config.AccessTokenSecret))
}

// Kontrola tokenu.
//
//	@param tokenStr
//	@param secret
//	@return uint
//	@return error
func TokenValid(tokenStr string, secret string) (uint, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("nečekána podepisovací metoda: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	fId := claims["sub"].(float64)
	return uint(fId), nil
}

// Vrátí token z kontextu
//
//	@param c
//	@return string
func ExtractTokenFromContext(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
