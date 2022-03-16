package model

import (
	"time"

	"github.com/beenson/URL_Shortener/pkg/repository"
	"github.com/beenson/URL_Shortener/service/database"
)

type Shorten struct {
	ID       uint      `gorm:"primaryKey;autoIncrement"`
	Code     string    `gorm:"not null"`
	URL      string    `gorm:"not null"`
	ExpireAt time.Time `gorm:"not null"`
}

func CreateShorten(shorten *Shorten) error {
	// Check if code is avaliable
	if !checkCodeAvaliable(shorten.Code) {
		return repository.ErrCodeUnavailable
	}

	// Insert into database
	if result := database.Instance.Create(&shorten); result.Error != nil {
		// Code conflict
		return result.Error
	}

	// Insert success
	return nil
}

func GetOriginalUrl(shorten *Shorten) error {
	if result := database.Instance.Model(&Shorten{}).Where("code = ? AND expire_at >= ?", shorten.Code, time.Now()).First(shorten); result.Error != nil {
		return result.Error
	}

	return nil
}

func checkCodeAvaliable(code string) bool {
	var count int64
	database.Instance.Model(&Shorten{}).Where("code = ? AND expire_at >= ?", code, time.Now()).Count(&count)
	return count == 0
}
