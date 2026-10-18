package auth

import (
	"fmt"

	core_errors "github.com/KyoshiBlame/RoshPit/event-api/internal/core/errors"
	"github.com/golang-jwt/jwt/v5"
)

func (m *Manager) ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpedted signing method: %v", t.Header["algs"])
			}
			return m.secret, nil
		})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, core_errors.ErrUnauthorized
	}

	return claims, nil
}
