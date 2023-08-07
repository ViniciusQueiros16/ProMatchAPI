package token

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func DecodeAuthToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method and return the secret key used for signing
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, fmt.Errorf("DecodeAuthToken: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userIDFloat, ok := claims["user_id"].(float64); ok {
			// Convert the float64 to int64
			userID := int64(userIDFloat)
			return userID, nil
		}
	}

	return 0, fmt.Errorf("Invalid token or user ID not found")
}
