package auth

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtSecretKey struct{}

func WithJwtSecretKey(ctx context.Context, secretKey []byte) context.Context {
	return context.WithValue(ctx, JwtSecretKey{}, secretKey)
}

func GetJwtSecretKey(ctx context.Context) ([]byte, error) {
	secretKey, ok := ctx.Value(JwtSecretKey{}).([]byte)
	if !ok {
		return nil, errors.New("jwt secret key not found")
	}

	return secretKey, nil
}

func AccessTokenLifetime() time.Duration {
	return 7 * 24 * time.Hour
}

func RefreshTokenLifetime() time.Duration {
	return 6 * 31 * 24 * time.Hour
}

func MakeUserToken(secretKey []byte, userId string, userNick string, lifetime time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"user_id":    userId,
		"user_nick":  userNick,
		"expired_at": time.Now().Add(lifetime),
	})

	value, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return value, nil
}

func ValidateUserToken(secretKey []byte, token string) (jwt.MapClaims, error) {
	claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS512 {
			return errors.New("invalid signed method"), nil
		}

		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !claims.Valid {
		return nil, errors.New("invalid token")
	}

	value, ok := claims.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return value, nil
}
