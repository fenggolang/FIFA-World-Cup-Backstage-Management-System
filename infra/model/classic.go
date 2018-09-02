package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 世界杯经典回顾
type WorldCupArchive struct {
	gorm.Model
	URL         string `gorm:"type:varchar(128);column:world_cup_url"`
	Name        string `gorm:"type:varchar(128);column:country_name"`
	Year        string `gorm:"type:varchar(64);column:year"`
	Winner      string `gorm:"type:varchar(64);column:winner_country"`
	RunnersUp   string `gorm:"type:varchar(64);column:runners_up_name"`
	Third       string `gorm:"type:varchar(64);column:third_name"`
	Fourth      string `gorm:"type:varchar(64);column:fourth_name"`
	FinalResult string `gorm:"type:varchar(128);column:final_result"`
	Title       string `gorm:"type:varchar(64);column:title"`
}

// 序列化为json
type WorldCupArchiveSerializer struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	URL         string     `json:"url"`
	Name        string     `json:"name"`
	Year        string     `json:"year"`
	Winner      string     `json:"winner"`
	RunnersUp   string     `json:"runners_up"`
	Third       string     `json:"third"`
	Fourth      string     `json:"fourth"`
	FinalResult string     `json:"final_result"`
	Title       string     `json:"title"`
}

func (w *WorldCupArchive) Serializer() WorldCupArchiveSerializer {
	return WorldCupArchiveSerializer{
		ID:          w.ID,
		CreatedAt:   w.CreatedAt,
		UpdatedAt:   w.UpdatedAt,
		DeletedAt:   w.DeletedAt,
		URL:         w.URL,
		Name:        w.Name,
		Year:        w.Year,
		Winner:      w.Winner,
		RunnersUp:   w.RunnersUp,
		Third:       w.Third,
		Fourth:      w.Fourth,
		FinalResult: w.FinalResult,
		Title:       w.Title,
	}
}
