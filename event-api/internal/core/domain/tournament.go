package domain

import "time"

type Tournament struct {
	ID                string
	Version           int64
	Name              string
	Description       string
	Status            TournamentStatus
	Format            TournamentFormat
	MaxTeams          int
	OrganizerID       string
	RegistrationStart time.Time
	RegistrationEnd   time.Time
	StartDate         time.Time
	EndDate           time.Time
	CreatedAt         time.Time
}

type TournamentStatus string

const (
	TournamentStatusDraft        TournamentStatus = "draft"
	TournamentStatusRegistration TournamentStatus = "registration"
	TournamentStatusInProgress   TournamentStatus = "in_progress"
	TournamentStatusFinished     TournamentStatus = "finished"
	TournamentStatusCancelled    TournamentStatus = "cancelled"
)

type TournamentFormat string

const (
	TournamentFormatOnline TournamentFormat = "online"
	TournamentFormatOflane TournamentFormat = "offline"
)
