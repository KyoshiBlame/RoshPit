package domain

import "time"

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type RefreshSession struct {
	TokenHash string
	UserID    string
	ExpiresAt time.Time
	CreatedAt time.Time
}
