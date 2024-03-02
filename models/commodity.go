package models

import (
	"github.com/opentreehole/go-common"
	"gorm.io/gorm"
)

type Commodity struct {
	ID             int    `json:"id" gorm:"primaryKey"`
	DefaultName    string `json:"default_name" gorm:"not null;size:64;index"`
	ProduceAt      MyTime `json:"produce_at" gorm:"not null"`
	ProduceAddress string `json:"produce_address" gorm:"not null;size:64"`
	Category       string `json:"category" gorm:"not null;size:64;index"`
}

func GetCommodities() (commodities []Commodity, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Find(&commodities).Error
	})
	return
}

func GetCommodityByID(commodityID int) (commodity *Commodity, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Take(&commodity, commodityID).Error
	})
	return
}

func (commodity *Commodity) Create() (err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&commodity).Error
	})
	return
}

func DeleteCommodityByID(commodityID int) (err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Delete(&Commodity{}, commodityID)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return common.NotFound("Commodity not found")
		}
		return nil
	})
	return
}

func (commodity *Commodity) Update() (err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Updates(&commodity)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return common.NotFound("Commodity not found")
		}
		return nil
	})
	return
}

func GetCommoditiesByName(name string) (commodities []Commodity, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Where("default_name = ?", name).Find(&commodities).Error
	})
	return
}

func GetCommoditiesByNameFuzzy(name string) (commodities []Commodity, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Where("default_name LIKE ?", "%"+name+"%").Find(&commodities).Error
	})
	return
}

func GetCommoditiesByCategory(category string) (commodities []Commodity, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Where("category = ?", category).Find(&commodities).Error
	})
	return
}

func GetCommoditiesByCategoryFuzzy(category string) (commodities []Commodity, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Where("category LIKE ?", "%"+category+"%").Find(&commodities).Error
	})
	return
}
