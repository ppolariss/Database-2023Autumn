package models

import (
	"github.com/opentreehole/go-common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommodityItem struct {
	ID          int        `json:"id" gorm:"primaryKey"`
	Commodity   *Commodity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommodityID int        `json:"commodity_id" gorm:"not null;uniqueIndex:idx_item"`
	Platform    *Platform  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PlatformID  int        `json:"platform_id" gorm:"not null;uniqueIndex:idx_item"`
	Seller      *Seller    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SellerID    int        `json:"seller_id" gorm:"not null;uniqueIndex:idx_item"`
	ItemName    string     `json:"item_name" gorm:"not null;size:64;index"`
	Price       float32    `json:"price" gorm:"check:price > 0; not null"`
	UpdateAt    MyTime     `json:"update_at" gorm:"autoUpdateTime;index:,sort:desc"`
}

func GetItemByID(ID int) (item *CommodityItem, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Take(&item, ID).Error
	})
	return
}

func GetItemsBySellerID(sellerID int) (items []CommodityItem, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Preload(clause.Associations).Where("seller_id = ?", sellerID).Find(&items).Error
	})
	return
}

func GetItems() (items []CommodityItem, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Preload(clause.Associations).Find(&items).Error
	})

	//for _, it := range items {
	//	//it.Commodity = new(Commodity)
	//	//it.Platform = new(Platform)
	//	//it.Seller = new(Seller)
	//	//_ = it.Commodity.GetCommodityByID(it.CommodityID)
	//	//_ = it.Platform.GetPlatformByID(it.PlatformID)
	//	//_ = it.Seller.GetSellerByID(it.SellerID)
	//	fmt.Println(it.CommodityID)
	//}
	return
}

func GetItemsByCommodityID(commodityID int) (items []CommodityItem, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Preload(clause.Associations).Where("commodity_id = ?", commodityID).Find(&items).Error
	})
	return
}

func (item *CommodityItem) Create() error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&item.Commodity, item.CommodityID).Error
		if err != nil {
			return err
		}
		err = tx.First(&item.Platform, item.PlatformID).Error
		if err != nil {
			return err
		}
		err = tx.First(&item.Seller, item.SellerID).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return common.NotFound("Commodity or Platform or Seller not found")
	}
	return DB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&item).Error
	})
}

func DeleteItemByID(itemID int) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&CommodityItem{}, itemID).Error
	})
}

func (item *CommodityItem) Update() error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		var err error
		if item.CommodityID != 0 {
			err = tx.First(&item.Commodity, item.CommodityID).Error
			if err != nil {
				return err
			}
		}
		if item.PlatformID != 0 {
			err = tx.First(&item.Platform, item.PlatformID).Error
			if err != nil {
				return err
			}
		}
		if item.SellerID != 0 {
			err = tx.First(&item.Seller, item.SellerID).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return common.NotFound("Commodity or Platform or Seller not found")
	}
	return DB.Transaction(func(tx *gorm.DB) error {
		//item.UpdateAt = MyTime{time.Now()}
		result := tx.Updates(&item)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			// when the value is not changed, it will return 0
			return common.NotFound("CommodityItem not found")
		}
		return nil
	})
}

func GetItemsByName(name string) (items []CommodityItem, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Preload(clause.Associations).Where("item_name = ?", name).Find(&items).Error
	})
	return
}

func GetItemsByNameFuzzy(name string) (items []CommodityItem, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.Preload(clause.Associations).Where("item_name LIKE ?", "%"+name+"%").Find(&items).Error
	})
	return
}

func GetItemsByCategory(category string) (items []CommodityItem, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.
			Joins("JOIN commodity ON commodity.id = commodity_item.commodity_id").
			Where("commodity.category = ?", category).
			Preload(clause.Associations).
			Find(&items).
			Error
	})
	return
}

func GetItemsByCategoryFuzzy(category string) (items []CommodityItem, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		return tx.
			Joins("JOIN commodity ON commodity.id = commodity_item.commodity_id").
			Where("commodity.category LIKE ?", "%"+category+"%").
			Preload(clause.Associations).
			Find(&items).
			Error
	})
	return
}

func CreateItems(items []CommodityItem) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&items).Error
	})
}

func (item *CommodityItem) AfterCreate(tx *gorm.DB) (err error) {
	var priceChange = PriceChange{
		CommodityItemID: item.ID,
		NewPrice:        item.Price,
		//UpdateAt:        MyTime{t},
	}
	//err = tx.Transaction(func(tx *gorm.DB) (err error) {
	// insert priceChange
	return tx.Create(&priceChange).Error
}

func (item *CommodityItem) AfterUpdate(tx *gorm.DB) (err error) {
	if item.Price == 0 {
		return
	}
	var favorites []Favorite
	var messages []Message
	//var t = time.Now()
	var priceChange = PriceChange{
		CommodityItemID: item.ID,
		NewPrice:        item.Price,
		//UpdateAt:        MyTime{t},
	}
	//err = tx.Transaction(func(tx *gorm.DB) (err error) {
	// insert priceChange
	err = tx.Create(&priceChange).Error
	if err != nil {
		return
	}
	// find favorites
	err = tx.Where("commodity_item_id = ? AND price_limit >= ?", item.ID, item.Price).Find(&favorites).Error
	if err != nil || len(favorites) == 0 {
		return
	}

	for _, favorite := range favorites {
		messages = append(messages, Message{
			UserID:          favorite.UserID,
			CommodityItemID: item.ID,
			CurrentPrice:    item.Price,
			//CreateAt:        MyTime{t},
		})
	}
	// insert messages
	return tx.Create(&messages).Error
	//})
}
