package models

import (
	"github.com/opentreehole/go-common"
	"gorm.io/gorm"
)

type Seller struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"uniqueIndex;not null;size:64"`
	Password string `json:"password" gorm:"not null;size:64"`
	Email    string `json:"email" gorm:"size:64"`
	//Age      int
	Address string `json:"address" gorm:"not null;size:64"`
}

//func (seller Seller) MarshalJSON() ([]byte, error) {
//	var sellerMap = map[string]interface{}{
//		"id":       seller.ID,
//		"username": seller.Username,
//		"email":    seller.Email,
//		"address":  seller.Address,
//	}
//	return json.Marshal(sellerMap)
//}

func GetSellers() (sellers []Seller, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Find(&sellers).Error
	})
	return
}

func GetSellerByID(userID int) (seller *Seller, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Omit("Password").Take(&seller, userID).Error
	})
	return
}

func (seller *Seller) Create() error {
	return DB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&seller).Error
	})
}

func (seller *Seller) Update() error {
	return DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Updates(&seller)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return common.NotFound("Seller not found")
		}
		return nil
	})
}

func DeleteSellerByID(id int) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		result := tx.Delete(&Seller{}, id)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return common.NotFound("Seller not found")
		}
		return nil
	})
}
