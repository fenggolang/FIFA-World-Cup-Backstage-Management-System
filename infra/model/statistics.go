package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// goals scored:进球数

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

// 统计球队----Attempts(试图)
/**
RANK: 排名
TEAM: 球队名称
MATCHES PLAYED: 个人上场比赛场次
SHOTS: 射门次数
ATTEMPTS ON TARGET: 射门射正次数
ATTEMPTS OFF TARGET: 射门射偏次数
SHOTS BLOCKED: 射门被挡出的次数
*/

// 统计球队---Disciplinary(犯规吃牌等情况)
/**
RANK: 排名
TEAM: 球队名称
MATCHES PLAYED: 球队比赛场次
YELLOW CARDS: 黄牌
INDIRECT RED CARDS: 间接红牌
DIRECT RED CARDS: 直接红牌
FOULS COMMITTED: 犯规
FOULS SUFFERED: 犯规遭受的
FOULS CAUSING PENALTY: 犯规导致的点球
*/

//============================================================================================================================================================
// https://www.fifa.com/worldcup/statistics/players/goal-scored
// 统计球员---Goals scored(进球数)
/**
RANK: 排名(依据个人总进球数排名，如果个人总进球数相同则依据个人总助攻数排名，如果还是一样，则依据个人总上场时间排名，上场时间短则排名靠前)
PLAYER: 球员姓名
GOALS SCORED：个人总进球数
ASSISTS: 个人总助攻数
MINUTES PLAYED: 个人上场跑动时间
MATCHES PLAYED: 个人上场比赛场次
PENALTIES SCORED: 通过点球进球个数
GOALS SCORED WITH THE LEFT FOOT：左脚进球的个数
GOALS SCORED WITH THE RIGHT FOOT：右脚进球的个数
HEADED GOALS：头球进球的个数
*/

// 统计球员---Top saves(守门员扑救[救球]总次数)
/**
RANK: 排名
PLAYER: 球员姓名
MATCHES PLAYED: 个人上场比赛场次
MINUTES PLAYED: 个人上场跑动时间
SAVES: 扑救次数
SAVE RATE: 扑救成功率
*/

// 统计球员---Shots(射门次数)
/**
RANK: 排名
PLAYER: 球员姓名
MATCHES PLAYED: 个人上场比赛场次
MINUTES PLAYED: 个人上场跑动时间
SHOTS: 射门次数
ATTEMPTS ON TARGET: 射门射正次数
ATTEMPTS OFF TARGET: 射门射偏次数
SHOTS BLOCKED: 射门被挡出的次数
*/

// 统计球员---Disciplinary(犯规吃牌等情况)
/**
RANK: 排名
PLAYER: 球员姓名
MATCHES PLAYED: 个人上场比赛场次
YELLOW CARDS: 黄牌
INDIRECT RED CARDS: 间接红牌
DIRECT RED CARDS: 直接红牌
FOULS COMMITTED: 犯规
FOULS SUFFERED: 犯规遭受的
FOULS CAUSING PENALTY: 犯规导致的点球
*/
