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

type JwtClaimsKey struct{}

func WithJwtClaims(ctx context.Context, claims jwt.MapClaims) context.Context {
	return context.WithValue(ctx, JwtClaimsKey{}, claims)
}

func GetJwtClaims(ctx context.Context) (jwt.MapClaims, error) {
	claims, ok := ctx.Value(JwtClaimsKey{}).(jwt.MapClaims)
	if !ok {
		return nil, errors.New("jwt claims not found")
	}

	return claims, nil
}

func AccessTokenLifetime() time.Duration {
	return 7 * 24 * time.Hour
}

func RefreshTokenLifetime() time.Duration {
	return 6 * 31 * 24 * time.Hour
}

const (
	UserId    = "user_id"
	UserNick  = "user_nick"
	ExpiredAt = "expired_at"
)

func MakeUserToken(secretKey []byte, userId string, userNick string, lifetime time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		UserId:    userId,
		UserNick:  userNick,
		ExpiredAt: time.Now().Add(lifetime).Unix(),
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

	expiredAt, ok := value[ExpiredAt].(float64)
	if !ok {
		return nil, errors.New("invalid expired at")
	}

	if time.Now().Unix() > int64(expiredAt) {
		return nil, errors.New("expired token")
	}

	return value, nil
}
