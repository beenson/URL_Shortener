package model

import "time"

type Shorten struct {
	Id       uint      `gorm:"primaryKey;autoIncrement"`
	Code     string    `gorm:"not null"`
	URL      string    `gorm:"not null"`
	ExpireAt time.Time `gorm:"not null"`
}
