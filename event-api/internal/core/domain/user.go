package domain

import "time"

type UserRole string

const (
	UserRoleViewer UserRole = "viewer"
	UserRoleJudge  UserRole = "judge"
	UserRoleAdmin  UserRole = "admin"
)

type User struct {
	ID           string
	Version      int64
	Email        string
	Nickname     string
	FirstName    string
	LastName     string
	Age          int
	Role         UserRole
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUser(
	email string,
	nickname string,
	firstName string,
	lastName string,
	age int,
	passwordHash string,
) User {
	now := time.Now().UTC()
	return User{
		Version:      1,
		Email:        email,
		Nickname:     nickname,
		FirstName:    firstName,
		LastName:     lastName,
		Age:          age,
		Role:         UserRoleViewer,
		PasswordHash: passwordHash,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

type PatchUser struct {
	Email     *string
	Nickname  *string
	FirstName *string
	LastName  *string
	Age       *int
}

type PublicUserProfile struct {
	ID        string    `json:"id"`
	Nickname  string    `json:"nickname"`
	Age       int       `json:"age"`
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
