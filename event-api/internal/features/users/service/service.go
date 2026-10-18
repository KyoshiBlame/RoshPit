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
	GetUserByID(
		ctx context.Context,
		id int,
	) (domain.PublicUserProfile, error)

	GetMe(
		ctx context.Context,
		id int,
	)

	PatchMe(
		ctx context.Context,
		id int,
	) (domain.User, error)

	DeleteMe(
		ctx context.Context,
		id int,
	) error
}
