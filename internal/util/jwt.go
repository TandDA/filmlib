package util

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/TandDA/filmlib/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

var privateKey = []byte("asf!E@@!KSF@!<Skq;zQweW18nNz") // TODO .env import

func GenerateJWT(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.Id,
		"role": user.RoleId,
		"iat":  time.Now().Unix(),
		"eat":  time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(privateKey)
}

func ValidateJWT(r *http.Request) error {
	token, err := getToken(r)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

// validate Admin role
func ValidateAdminRoleJWT(r *http.Request) error {
	token, err := getToken(r)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 1 {
		return nil
	}
	return errors.New("invalid admin token provided")
}

// check token validity
func getToken(r *http.Request) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}