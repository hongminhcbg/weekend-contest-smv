package models

import "time"

type UserVisited struct {
	Id       int64  `json:"id" gorm:"type:SERIAL;primary_key;AUTO_INCREMENT"`
	Ip       string `json:"ip" gorm:"column:ip"`
	Browser  string `json:"browser" gorm:"column:browser"`
	Location string `json:"location" gorm:"column:location"`

	CreatedAt *time.Time `json:"-" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"-" gorm:"column:updated_at"`
}

func (UserVisited) TableName() string {
	return "user_visited"
}
