package domain

import "time"

type TournamentJudge struct {
	TournamentID string
	UserID       string
	AssignedAT   time.Time
}


