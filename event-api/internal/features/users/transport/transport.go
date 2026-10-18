package transport

import (
	"context"

	"github.com/KyoshiBlame/RoshPit/event-api/internal/core/domain"
	"github.com/gin-gonic/gin"
)

type UsersHTTPHandler struct {
	UserService UsersService
}

func NewUsersHTTPHandler(
	userService UsersService,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		UserService: userService,
	}
}

type UsersService interface {
	GetMe(
		ctx context.Context,
		id int,
	) (domain.User, error)

	PatchMe(
		ctx context.Context,
		id int,
		patch domain.PatchUser,
	) (domain.User, error)

	GetUserByID(
		ctx context.Context,
		id int,
	) (domain.PublicUserProfile, error)

	DeleteMe(
		ctx context.Context,
		id int,
	) error
}

func (h *UsersHTTPHandler) RegisterRoutes(
	router gin.IRouter,
	authMiddleWare gin.HandlerFunc,
) {
	usersGroup := router.Group("/users")
	{
		// Публичный роут (доступен всем без токена)
		usersGroup.GET("/:id", h.GetUserByID)

		// Защищенные роуты (только с валидным JWT)
		protected := usersGroup.Group("")
		protected.Use(authMiddleWare)
		{
			protected.GET("/me", h.GetMe)
			protected.PATCH("/me", h.PatchMe)
			protected.DELETE("/me", h.DeleteMe)
			// Сюда легко добавлять другие закрытые методы:
			// protected.PATCH("/me", h.UpdateMe)
		}
	}
}
