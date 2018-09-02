package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 球员信息
type Player struct {
	gorm.Model
	Number   int    `gorm:"type:integer;not null;column:number"`             //球员球衣号
	Name     string `gorm:"type:varchar(64);not null;column:name"`           //球员姓名
	Country  string `gorm:"type:varchar(64);not null;column:country"`        //球员国家
	Role     string `gorm:"type:varchar(64);not null;column:role"`           //球员角色(中场，前锋、后场、、、)
	ImageURL string `gorm:"type:varchar(128);not null;column:image_address"` //球员个人图片
}
type PlayerSerializer struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Number    int        `json:"number"`
	Name      string     `json:"name"`
	Country   string     `json:"country"`
	Role      string     `json:"role"`
	ImageURL  string     `json:"image_url"`
}

func (p *Player) Serializer() PlayerSerializer {
	return PlayerSerializer{
		ID:        p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		DeletedAt: p.DeletedAt,
		Number:    p.Number,
		Name:      p.Name,
		Country:   p.Country,
		Role:      p.Role,
		ImageURL:  p.ImageURL,
	}
}

// 教练信息
type Coach struct {
	gorm.Model
	CountryName string `gorm:"type:varchar(64);not null;column:country_name"`
	Name        string `gorm:"type:varchar(64);not null;column:name"`
	ImageURL    string `gorm:"type:varchar(128);not null;column:image_address"`
}

type CoachSerializer struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	CountryName string     `json:"country_name"`
	Name        string     `json:"name"`
	ImageURL    string     `json:"image_url"`
}

func (c *Coach) Serializer() CoachSerializer {
	return CoachSerializer{
		ID:          c.ID,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
		DeletedAt:   c.DeletedAt,
		CountryName: c.CountryName,
		Name:        c.Name,
		ImageURL:    c.ImageURL,
	}
}
