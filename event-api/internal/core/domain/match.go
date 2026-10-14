package domain

import "time"

type Match struct {
	ID           string
	Version      int64
	TournamentID string
	TeamAID      string
	TeamBID      string
	JudgeID      string
	Status       MatchStatus
	Round        int
	ScheduledAt  time.Time
	StartedAt    *time.Time // nil пока не начался
	FinishedAt   *time.Time // nil пока не завершён
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type MatchStatus string

const (
	MatchStatusScheduled  MatchStatus = "scheduled"
	MatchStatusInProgress MatchStatus = "in_progress"
	MatchStatusPaused     MatchStatus = "paused"
	MatchStatusFinished   MatchStatus = "finished"
	MatchStatusCancelled  MatchStatus = "cancelled"
)

type MatchScore struct {
	MatchID  string
	TeamID   string
	Version  int64
	Score    int
	UpdateAt time.Time
}
