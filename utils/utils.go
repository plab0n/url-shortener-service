package utils

import (
	"fmt"
	"strings"
	"url-shortener-service/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func ParseToken(tokenString string) (claims *models.Claims, err error) {
	fmt.Println(tokenString)
	tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwtKey")), nil
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	claims, ok := token.Claims.(*models.Claims)

	if !ok {
		return nil, err
	}

	return claims, nil
}
