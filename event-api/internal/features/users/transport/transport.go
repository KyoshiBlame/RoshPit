package transport

import (
	"context"

	"github.com/KyoshiBlame/RoshPit/event-api/internal/core/domain"
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
	GetUser(
		ctx context.Context,
		nickname string,
	) (domain.User, error)

	UpdateUser(
		ctx context.Context,
		id int,
		patch domain.PatchUser,
	) (domain.User, error)

	GetPublicUserByID(
		ctx context.Context,
		id int,
	) (domain.PublicUserProfile, error)

	DeleteUser(
		ctx context.Context,
		id int,
	) error
}
