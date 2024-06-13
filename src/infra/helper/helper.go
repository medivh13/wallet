package helper

import (
	"errors"
	"time"
	dto "wallet/src/app/dto/user"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("secret_key")

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrUserNotFound        = errors.New("destination user not found")
)

// TokenClaims menyimpan klaim JWT
type TokenClaims struct {
	UserID   int64  `json:"user_id"`
	WalletID int64  `json:"wallet_id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken membuat token JWT
func GenerateToken(data *dto.RegisterModel) (string, error) {
	expirationTime := time.Now().Add(120 * time.Minute)
	claims := &TokenClaims{
		UserID:   data.ID,
		WalletID: data.WalletID,
		UserName: data.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// VerifyToken memverifikasi token JWT
func VerifyToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
