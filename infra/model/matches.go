package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 比赛
type Match struct {
	gorm.Model
	Date         string `gorm:"type:varchar(64);not null;column:date"`
	GroupName    string `gorm:"type:varchar(64);not null;column:group_name"`
	Location     string `gorm:"type:varchar(64);not null;column:location"`
	CountryLeft  string `gorm:"type:varchar(64);not null;column:country_left"`
	CountryRight string `gorm:"type:varchar(64);not null;column:country_right"`
	Score        string `gorm:"type:varchar(64);not null;column:score"`
	MatchNumber  int    `gorm:"type:integer;not null;column:match_number"`
}

func (m Match) TableName() string {
	return "matches"
}

type MatchSerializer struct {
	ID           uint       `json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Date         string     `json:"date"`
	GroupName    string     `json:"group_name"`
	Location     string     `json:"location"`
	CountryLeft  string     `json:"country_left"`
	CountryRight string     `json:"country_right"`
	Score        string     `json:"score"`
	MatchNumber  int        `json:"match_number"`
}

func (m Match) Serializer() MatchSerializer {
	return MatchSerializer{
		ID:           m.ID,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
		DeletedAt:    m.DeletedAt,
		Date:         m.Date,
		GroupName:    m.GroupName,
		Location:     m.Location,
		CountryLeft:  m.CountryLeft,
		CountryRight: m.CountryRight,
		Score:        m.Score,
		MatchNumber:  m.MatchNumber,
	}
}
