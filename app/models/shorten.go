package model

import (
	"encoding/json"
	"log"
	"time"

	"github.com/beenson/URL_Shortener/pkg/repository"
	"github.com/beenson/URL_Shortener/service/cache"
	"github.com/beenson/URL_Shortener/service/database"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Shorten struct {
	ID       uint      `gorm:"primaryKey;autoIncrement"`
	Code     string    `gorm:"not null"`
	URL      string    `gorm:"not null"`
	ExpireAt time.Time `gorm:"not null"`
}

func (shorten *Shorten) AfterSave(tx *gorm.DB) (err error) {
	if !checkCodeAvailable(shorten.Code) {
		return repository.ErrCodeUnavailable
	}
	return
}

func CreateShorten(shorten *Shorten) error {
	// Insert into database
	if result := database.Instance.Create(&shorten); result.Error != nil {
		// Code conflict
		return result.Error
	}

	// Insert success
	return nil
}

func GetOriginalUrl(shorten *Shorten) error {
	// check if exist in redis first
	if getInfoFromCache(shorten.Code, shorten) {
		// if expired return error
		if shorten.ExpireAt.Before(time.Now()) {
			return gorm.ErrRecordNotFound
		}
		return nil
	}

	if result := database.Instance.Model(&Shorten{}).Where("code = ? AND expire_at >= ?", shorten.Code, time.Now()).First(shorten); result.Error != nil {
		return result.Error
	}

	// write back to redis
	if err := writeInfoIntoCache(shorten); err != nil {
		log.Fatal(err.Error())
		return nil
	}

	return nil
}

// true for available; false for unavailable
func checkCodeAvailable(code string) bool {
	if getInfoFromCache(code, nil) {
		return false
	}

	var count int64
	database.Instance.Model(&Shorten{}).Where("code = ? AND expire_at >= ?", code, time.Now()).Count(&count)
	return count == 0
}

// ture for found; false for not found
func getInfoFromCache(code string, shorten *Shorten) bool {
	val, err := cache.Instance.Get(cache.Ctx, code).Result()
	if err == redis.Nil {
		// code not exist in redis
		return false
	} else if err != nil {
		log.Fatal(err.Error())
		return false
	}

	// Don't need entity
	if shorten == nil {
		return true
	}

	err = json.Unmarshal([]byte(val), shorten)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}

	log.Printf("load %s from cache", shorten.Code)
	return true
}

func writeInfoIntoCache(shorten *Shorten) error {
	b, err := json.Marshal(shorten)
	if err != nil {
		return err
	}

	if err := cache.Instance.Set(cache.Ctx, shorten.Code, b, 3600*time.Second).Err(); err != nil {
		return err
	}

	log.Printf("cache %s", shorten.Code)
	return nil
}
