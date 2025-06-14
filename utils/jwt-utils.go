package utils

import (
	"errors"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenClaims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	jwt.RegisteredClaims
}

func ExtractAccessToken(c *gin.Context) (string, error) {
	// First try to get token from cookie
	if token, err := c.Cookie("access_token"); err == nil && token != "" {
		return token, nil
	}

	return "", errors.New("no token found")
}

func ExtractRefreshToken(c *gin.Context) (string, error) {
	token, err := c.Cookie("refresh_token")
	if err == nil && token != "" {
		return token, nil
	}

	return "", errors.New("no refresh token found")
}

func GenerateTokens(userID uuid.UUID, email string) (string, string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", "", errors.New("JWT_SECRET not set")
	}

	accessDuration, _ := time.ParseDuration(os.Getenv("JWT_ACCESS_TOKEN_EXPIRY"))
	refreshDuration, _ := time.ParseDuration(os.Getenv("JWT_REFRESH_TOKEN_EXPIRY"))

	now := time.Now()
	claims := func(duration time.Duration) *TokenClaims {
		return &TokenClaims{
			UserID: userID,
			Email:  email,
			RegisteredClaims: jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(now),
				ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
				Issuer:    "OnePolicy",
			},
		}
	}

	signToken := func(claims *TokenClaims) (string, error) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString([]byte(secret))
	}

	accessToken, err := signToken(claims(accessDuration))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := signToken(claims(refreshDuration))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func ValidateToken(tokenString string) (*TokenClaims, error) {

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET not set")
	}

	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token.Claims.(*TokenClaims), nil
}
