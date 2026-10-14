package domain

import "time"

type User struct {
	ID           string
	Version      int64
	Email        string
	Nickname     string
	FirstName    string
	LastName     string
	Age          int
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUser(
	email string,
	nickname string,
	firstName string,
	LastName string,
	Age string,
)

type PatchUser struct {
	Email    string
	Nickname string
	LastName string
	Age      string
}

type PublicUserProfile struct {
	ID        string
	Nickname  string
	Age       string
	CreatedAt time.Time
}
