package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Group struct {
	gorm.Model
	GroupName   string `gorm:"type:varchar(64);not null;column:group_name"`
	TeamName    string `gorm:"type:varchar(64);not null;column:team_name"`
	MatchPlayed int    `gorm:"type:integer;not null;column:match_played"`
	WinNumber   int    `gorm:"type:integer;not null;column:win_number"`  // 赢局
	Draw        int    `gorm:"type:integer;not null;column:draw_number"` //平局
	Lost        int    `gorm:"type:integer;not null;column:lost_number"` //输局
	GoalFor     int    `gorm:"type:integer;not null;column:goal_number"`
	GoalAgainst int    `gorm:"type:integer;not null;column:goal_against"`
	DiffGoal    int    `gorm:"type:integer;not null;column:diff_goal"`
	Points      int    `gorm:"type:integer;not null;column:points"`
}

type GroupSerializer struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	GroupName   string     `json:"group_name"`
	TeamName    string     `json:"team_name"`
	MatchPlayed int        `json:"match_played"`
	WinNumber   int        `json:"win_number"`
	Draw        int        `json:"draw_number"`
	Lost        int        `json:"lost_number"`
	GoalFor     int        `json:"goal_number"`
	GoalAgainst int        `json:"goal_against"`
	DiffGoal    int        `json:"diff_goal"`
	Points      int        `json:"points"`
}

func (g *Group) Serializer() GroupSerializer {
	return GroupSerializer{
		ID:          g.ID,
		CreatedAt:   g.CreatedAt,
		UpdatedAt:   g.UpdatedAt,
		DeletedAt:   g.DeletedAt,
		GroupName:   g.GroupName,
		TeamName:    g.TeamName,
		MatchPlayed: g.MatchPlayed,
		WinNumber:   g.WinNumber,
		Draw:        g.Draw,
		Lost:        g.Lost,
		GoalFor:     g.GoalFor,
		GoalAgainst: g.GoalAgainst,
		DiffGoal:    g.DiffGoal,
		Points:      g.Points,
	}
}
