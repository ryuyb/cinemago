package utils

import (
	"cinemago/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type JwtClaims struct {
	jwt.RegisteredClaims
}

func GenerateAccessToken(id int) (string, error) {
	jwtCfg := config.GetConfig().Jwt

	expiresAt := time.Now().Add(time.Minute * time.Duration(jwtCfg.ValidWithinMinutes))
	uid, _ := uuid.NewV7()
	claims := JwtClaims{
		jwt.RegisteredClaims{
			Subject:   strconv.Itoa(id),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uid.String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtCfg.SigningKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
