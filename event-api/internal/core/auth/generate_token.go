package auth

import (
	"time"

	"github.com/KyoshiBlame/RoshPit/event-api/internal/core/domain"
	"github.com/golang-jwt/jwt/v5"
)

func (m *Manager) GenerateToken(userID string, role domain.UserRole) (string, error) {
	now := time.Now().UTC()

	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(m.accessTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.secret)
}
