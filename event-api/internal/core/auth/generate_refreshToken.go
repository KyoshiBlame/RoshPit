package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func (m *Manager) GenerateRegresToken() (rawToken string, tokenHash string, err error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", "", fmt.Errorf("read random bytes: %w", err)
	}

	rawToken = base64.URLEncoding.EncodeToString(b)
	tokenHash = HashToken(rawToken)
	return rawToken, tokenHash, nil
}

func HashToken(rawToken string) string {
	sum := sha256.Sum256([]byte(rawToken))
	return hex.EncodeToString(sum[:])
}
