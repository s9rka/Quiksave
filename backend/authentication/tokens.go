package auth

import (
	"encoding/hex"
	"crypto/rand"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID string `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)

	userIDStr := strconv.Itoa(userID)

	claims := &Claims{
		UserID: userIDStr,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "nota_bene",
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken.SignedString(jwtSecret)
}

func generateJTI() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func GenerateRefreshToken(userID int) (string, error) {
	jti, err := generateJTI()
	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(24 * time.Hour)

	// Claims with UserID, jti, and expiry
	claims := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(userID),
		ID:        jti,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		Issuer:    "nota_bene",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return refreshToken.SignedString(jwtSecret)
}

func ValidateJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
