package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// https://www.fifa.com/worldcup/statistics/teams/goal-scored
// 统计球队－--Top goals
type TeamStatisticWithTopGoal struct {
	gorm.Model           //
	Rank          int    `gorm:"type:integer;not null;column:rank"`           //排名
	TeamName      string `gorm:"type:varchar(12);not null;column:team_name"`  //球队名
	MatchPlayed   int    `gorm:"type:integer;not null;column:matches_played"` //球队打的比赛场次
	GoalsFor      int    `gorm:"type:integer;not null;column:goals_for"`      //球队总进球数(包括加时进球数)
	GoalsScored   int    `gorm:"type:integer;not null;column:goals_scored"`   //
	GoalsAgainst  int    `gorm:"type:integer;not null;column:goals_against"`  //球队总的失球数(包括加时赛丢球数)
	PenaltyGoal   int    `gorm:"type:integer;not null;column:penalty_goal"`   //罚球目标
	OwnGoalsFor   int    `gorm:"type:integer;not null;column:own_goals_for"`  //自己的目标
	OpenPlayGoals int    `gorm:"type:integer;not null;column:open_play_goals"`
	SetPieceGoals int    `gorm:"type:integer;not null;column:set_piece_goals"`
}
type TeamStatisticWithTopGoalSerializer struct {
	ID            uint       `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	Rank          int        `json:"rank"`
	TeamName      string     `json:"team_name"`
	MatchPlayed   int        `json:"match_played"`
	GoalsFor      int        `json:"goals_for"`
	GoalsScored   int        `json:"goals_scored"`
	GoalsAgainst  int        `json:"goals_against"`
	PenaltyGoal   int        `json:"penalty_goal"`
	OwnGoalsFor   int        `json:"own_goals_for"`
	OpenPlayGoals int        `json:"open_play_goals"`
	SetPieceGoals int        `json:"set_piece_goals"`
}

func (ttg *TeamStatisticWithTopGoal) Serializer() TeamStatisticWithTopGoalSerializer {
	return TeamStatisticWithTopGoalSerializer{
		ID:            ttg.ID,
		CreatedAt:     ttg.CreatedAt,
		UpdatedAt:     ttg.UpdatedAt,
		DeletedAt:     ttg.DeletedAt,
		Rank:          ttg.Rank,
		TeamName:      ttg.TeamName,
		MatchPlayed:   ttg.MatchPlayed,
		GoalsFor:      ttg.GoalsFor,
		GoalsScored:   ttg.GoalsScored,
		GoalsAgainst:  ttg.GoalsAgainst,
		PenaltyGoal:   ttg.PenaltyGoal,
		OwnGoalsFor:   ttg.OwnGoalsFor,
		SetPieceGoals: ttg.SetPieceGoals,
	}
}

// 统计球队----Attempts

// 统计球队---Disciplinary

//=====================================================================================================
// https://www.fifa.com/worldcup/statistics/players/goal-scored
// 统计球员---Goals scored

// 统计球员---Top saves

// 统计球员---Shots

// 统计球员---Disciplinary
