package middleware

import (
	"net/http"
	"strings"

	"github.com/KyoshiBlame/RoshPit/event-api/internal/core/auth"
	"github.com/KyoshiBlame/RoshPit/event-api/internal/core/domain"
	"github.com/gin-gonic/gin"
)

const (
	ctxUserIDKey   = "user_id"
	ctxUserRoleKey = "user_role"
)

type TokenParser interface {
	ParseToken(tokenStr string) (*auth.Claims, error)
}

func Auth(parser TokenParser) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid authorization header"})
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")
		claims, err := parser.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		c.Set(ctxUserIDKey, claims.UserID)
		c.Set(ctxUserRoleKey, claims.Role)
		c.Next()
	}
}

func RequireRoles(allowedRoles ...domain.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, exists := c.Get(ctxUserIDKey)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		}

		userRole, ok := val.(domain.UserRole)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid role in context"})
			return
		}

		for _, role := range allowedRoles {
			if userRole == role {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
	}
}

func GetUserID(c *gin.Context) string {
	return c.GetString(ctxUserIDKey)
}

func GetUSerRoke(c *gin.Context) domain.UserRole {
	val, exists := c.Get(ctxUserRoleKey)
	if !exists {
		return ""
	}
	role, _ := val.(domain.UserRole)
	return role
}
