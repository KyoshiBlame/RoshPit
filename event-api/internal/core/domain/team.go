package domain

import "time"

type Team struct {
	ID           string
	Version      int64
	TournamentID string
	Name         string
	CaptainID    string
	Members      []TeamMember
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type TeamMember struct {
	TeamID   string
	UserID   string
	Role     TeamMemberRole
	JoinedAt time.Time
}

type TeamMemberRole string

const (
	TeamMemberCaptain TeamMemberRole = "captain"
	TeamMemberPlayer  TeamMemberRole = "player"
)
