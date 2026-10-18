package auth

import (
	"time"

	"github.com/KyoshiBlame/RoshPit/event-api/internal/core/domain"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string          `json:"user_id"`
	Role   domain.UserRole `json:"role"`
	jwt.RegisteredClaims
}

type Manager struct {
	secret     []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewManager(cfg Config) *Manager {
	return &Manager{
		secret:     []byte(cfg.Secret),
		accessTTL:  cfg.AccessTTL,
		refreshTTL: cfg.RefreshTTL,
	}
}

func (m *Manager) RefreshTTL() time.Duration {
	return m.refreshTTL
}
