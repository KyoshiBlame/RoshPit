package service

import (
	"context"

	"github.com/KyoshiBlame/RoshPit/event-api/internal/core/domain"
)

type UsersService struct {
	UserRepository UsersRepository
}

func NewUsersService(
	usersRepository UsersRepository,
) *UsersService {
	return &UsersService{
		UserRepository: usersRepository,
	}
}

type UsersRepository interface {
	GetUser(
		ctx context.Context,
		id int,
	) (domain.User, error)

	PatchUser(
		ctx context.Context,
		id int,
		patchUser domain.PatchUser,
	) (domain.User, error)
}
