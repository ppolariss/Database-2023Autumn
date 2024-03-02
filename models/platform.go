package models

import (
	"github.com/opentreehole/go-common"
	"gorm.io/gorm"
)

type Platform struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"not null;size:64"`
	URL     string `json:"url" gorm:"size:64"`
	Country string `json:"country" gorm:"size:64"`
}

func GetPlatforms() (platforms []Platform, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Find(&platforms).Error
	})
	return
}

func GetPlatformByID(platformID int) (platform *Platform, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Take(&platform, platformID).Error
	})
	return
}

func (platform *Platform) Create() error {
	return DB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&platform).Error
	})
}

func (platform *Platform) Update() error {
	return DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Updates(&platform)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return common.NotFound("Platform not found")
		}
		return nil
	})
}

func DeletePlatformByID(platformID int) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Delete(&Platform{}, platformID)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return common.NotFound("Platform not found")
		}
		return nil
	})
}
