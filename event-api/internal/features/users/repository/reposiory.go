package repository

import "github.com/KyoshiBlame/RoshPit/event-api/internal/core/repository"

type UsersRepository struct {
	pool repository.Pool
}

func NewUsersRepository(
	pool repository.Pool,
) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}
