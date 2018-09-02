package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Team struct {
	gorm.Model
	Flag string `gorm:"type:varchar(128);not null;column:flag_address"`
	Name string `gorm:"type:varchar(64);not null;column:team_name"`
}

type TeamSerializer struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Flag      string     `json:"flag"`
	Name      string     `json:"name"`
}

func (t *Team) Serializer() TeamSerializer {
	return TeamSerializer{
		ID:        t.ID,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		DeletedAt: t.DeletedAt,
		Flag:      t.Flag,
		Name:      t.Name,
	}
}
