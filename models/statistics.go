package models

import (
	"github.com/opentreehole/go-common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func GetFavoriteStatisticsAll() (favoriteStatisticsResponse []FavoriteStatisticsResponse, err error) {
	err = DB.Transaction(func(tx *gorm.DB) (err error) {
		var favStats []FavoriteStatistics
		err = tx.Model(&Favorite{}).
			Limit(10).
			Select("count(*) as count, commodity_item_id").
			Group("commodity_item_id").
			Order("count desc").
			Scan(&favStats).Error
		if err != nil {
			return
		}
		for _, stat := range favStats {
			var commodityItem CommodityItem
			err = tx.
				Preload(clause.Associations).
				First(&commodityItem, stat.CommodityItemID).Error
			if err != nil {
				return
			}
			favoriteStatisticsResponse = append(favoriteStatisticsResponse, FavoriteStatisticsResponse{
				CommodityItem: &commodityItem,
				Count:         stat.Count,
			})
		}
		return
	})
	return
	// RAW sql statement:
	// SELECT COUNT(*) AS COUNT, commodity_item_id FROM favorite GROUP BY commodity_item_id ORDER BY COUNT DESC LIMIT 20;
}

func GetFavoriteStatisticsSome(gender *bool, ageStart int, ageEnd int) (favoriteStatisticsResponse []FavoriteStatisticsResponse, err error) {
	err = DB.Transaction(func(tx *gorm.DB) (err error) {
		var favStats []FavoriteStatistics
		if gender != nil {
			err = tx.Model(&Favorite{}).
				Joins("JOIN user ON user.id = favorite.user_id").
				Where("user.gender = ? AND user.age BETWEEN ? AND ?", gender, ageStart, ageEnd).
				Limit(10).
				Select("count(*) as count, commodity_item_id").
				Group("commodity_item_id").
				Order("count DESC").
				Scan(&favStats).Error
		} else {
			err = tx.Model(&Favorite{}).
				Joins("JOIN user ON user.id = favorite.user_id").
				Where("user.age BETWEEN ? AND ?", ageStart, ageEnd).
				Limit(10).
				Select("count(*) as count, commodity_item_id").
				Group("commodity_item_id").
				Order("count DESC").
				Scan(&favStats).Error
		}
		if err != nil {
			return
		}
		for _, stat := range favStats {
			var commodityItem CommodityItem
			err = tx.
				Preload(clause.Associations).
				First(&commodityItem, stat.CommodityItemID).Error
			if err != nil {
				return
			}
			favoriteStatisticsResponse = append(favoriteStatisticsResponse, FavoriteStatisticsResponse{
				CommodityItem: &commodityItem,
				Count:         stat.Count,
			})
		}
		return
	})
	return
	// RAW sql statement:
	// SELECT COUNT(*) AS COUNT, commodity_item_id FROM favorite INNER JOIN USER ON favorite.`user_id` = user.`id` WHERE user.`gender` = FALSE AND user.`age` BETWEEN 10 AND 50 GROUP BY commodity_item_id ORDER BY COUNT DESC LIMIT 20;
	// SELECT COUNT(*) AS COUNT, commodity_item_id FROM favorite LEFT JOIN USER ON favorite.`user_id` = user.`id` WHERE user.`age` BETWEEN 10 AND 50 GROUP BY commodity_item_id ORDER BY COUNT DESC LIMIT 20;

	// Preload("CommodityItem").
	//	preload will happen after the query, so it will not work if select CommodityItem
}

func GetPriceStatisticsHistory(commodityID int, timeStart time.Time, timeEnd time.Time) (priceStatisticsResponse []PriceStatisticsResponse, err error) {
	err = DB.Transaction(func(tx *gorm.DB) (err error) {
		var items []CommodityItem
		err = tx.Model(&CommodityItem{}).
			Where("commodity_id = ?", commodityID).
			Preload(clause.Associations).
			Find(&items).Error
		if err != nil {
			return
		}
		if len(items) == 0 {
			return common.NotFound("Commodity not found")
		}
		for _, item := range items {
			var PriceVariance float32
			err = tx.Model(&PriceChange{}).
				Where("commodity_item_id = ? AND update_at BETWEEN ? AND ?", item.ID, timeStart, timeEnd).
				Group("commodity_item_id").
				Select("MAX(new_price) - MIN(new_price) as price_variance").
				Find(&PriceVariance).Error
			// SELECT MAX(new_price) - MIN(new_price) AS price_variance FROM price_change WHERE commodity_item_id = 1 GROUP BY commodity_item_id
			if err != nil {
				return
			}
			priceStatisticsResponse = append(priceStatisticsResponse, PriceStatisticsResponse{
				CommodityItem: item,
				PriceVariance: PriceVariance,
			})
		}
		return
	})
	// if transfer into a sql statement, it will be:
	// SELECT commodity_item_id, MAX(new_price) - MIN(new_price) AS price_variance FROM price_change WHERE commodity_item_id IN (SELECT id FROM commodity_item WHERE commodity_id = 1) AND update_at BETWEEN '2021-05-01 00:00:00' AND '2025-05-31 00:00:00' GROUP BY commodity_item_id
	// select MAX(new_price) - MIN(new_price) as price_variance, commodity_item_id from commodity_item join price_change on commodity_item.id = price_change.commodity_item_id where commodity_id = 54 AND price_change.update_at BETWEEN "2022-01-01 09:26:50.000" AND "2024-01-01 09:26:50.000" Group By commodity_item_id
	return
}

func GetAnnualSummary(userID int) (res SummaryResponse, err error) {
	//year, err := time.Parse("2006-01-02 15:04:05", "2023-01-01 00:00:00")
	//if err != nil {
	//	return
	//}
	err = DB.Transaction(func(tx *gorm.DB) (err error) {
		// favoriteNum
		var favoriteNum int64
		if userID != -1 {
			err = tx.Model(&Favorite{}).
				Where("user_id = ?", userID).
				Count(&favoriteNum).Error
		} else {
			err = tx.Model(&Favorite{}).
				Count(&favoriteNum).Error
		}
		if err != nil {
			return
		}
		// messageNum
		var messageNum int64
		if userID != -1 {
			err = tx.Model(&Message{}).
				Where("user_id = ?", userID).
				Count(&messageNum).Error
		} else {
			err = tx.Model(&Message{}).
				Count(&messageNum).Error
		}
		if err != nil {
			return
		}
		// priceChangeNum
		var priceChangeNum int64
		if userID != -1 {
			err = tx.Model(&Favorite{}).
				Joins("left join price_change on price_change.commodity_item_id = favorite.commodity_item_id").
				Where("favorite.user_id = ? AND favorite.update_at < price_change.update_at", userID).
				Count(&priceChangeNum).Error
			//err = tx.Raw("select count(*) from favorite left join price_change on price_change.commodity_item_id = favorite.commodity_item_id where favorite.user_id = ? AND favorite.update_at < price_change.update_at", userID).Scan(&priceChangeNum).Error
			// select count(*) from favorite left join price_change on price_change.commodity_item_id = favorite.commodity_item_id where favorite.user_id = 1 AND favorite.update_at < price_change.update_at
			// Get all:
			// SELECT COUNT(*) AS num,user_id FROM favorite LEFT JOIN price_change ON price_change.commodity_item_id = favorite.commodity_item_id WHERE favorite.update_at < price_change.update_at GROUP BY user_id ORDER BY num DESC
		} else {
			err = tx.Model(&Favorite{}).
				Joins("left join price_change on price_change.commodity_item_id = favorite.commodity_item_id").
				Where("favorite.update_at < price_change.update_at").
				Count(&priceChangeNum).Error
		}
		if err != nil {
			return
		}

		// commodity
		c := struct {
			CommodityID  int       `json:"commodity_id"`
			CommodityNum int       `json:"commodity_num"`
			Commodity    Commodity `json:"commodity"`
		}{}
		if userID != -1 {
			err = tx.Model(&Favorite{}).
				Where("user_id = ?", userID).
				Joins("left JOIN commodity_item ON commodity_item.id = favorite.commodity_item_id").
				Select("count(*) as commodity_num, commodity_id").
				Group("commodity_id").
				Order("commodity_num DESC").
				Limit(1).
				Scan(&c).Error
			// SELECT COUNT(*) AS commodity_num, commodity_id FROM favorite LEFT JOIN commodity_item ON commodity_item.id = favorite.commodity_item_id WHERE user_id = 2 GROUP BY commodity_id ORDER BY commodity_num DESC
		} else {
			err = tx.Model(&Favorite{}).
				Joins("left JOIN commodity_item ON commodity_item.id = favorite.commodity_item_id").
				Select("count(*) as commodity_num, commodity_id").
				Group("commodity_id").
				Order("commodity_num DESC").
				Limit(1).
				Scan(&c).Error
		}
		if err != nil {
			return
		}
		err = tx.Find(&c.Commodity, c.CommodityID).Error
		if err != nil {
			return
		}

		// seller
		var s = struct {
			SellerID  int    `json:"seller_id"`
			SellerNum int    `json:"seller_num"`
			Seller    Seller `json:"seller"`
		}{}
		if userID != -1 {
			err = tx.Model(&Favorite{}).
				Joins("left JOIN commodity_item ON commodity_item.id = favorite.commodity_item_id").
				Where("user_id = ?", userID).
				Select("count(*) as seller_num, seller_id").
				Group("seller_id").
				Order("seller_num DESC").
				Limit(1).
				Scan(&s).Error
			// SELECT COUNT(*) AS seller_num, seller_id FROM favorite LEFT JOIN commodity_item ON commodity_item.id = favorite.commodity_item_id WHERE user_id = 2 GROUP BY seller_id ORDER BY seller_num DESC
		} else {
			err = tx.Model(&Favorite{}).
				Joins("left JOIN commodity_item ON commodity_item.id = favorite.commodity_item_id").
				Select("count(*) as seller_num, seller_id").
				Group("seller_id").
				Order("seller_num DESC").
				Limit(1).
				Scan(&s).Error
		}
		if err != nil {
			return
		}
		err = tx.Find(&s.Seller, s.SellerID).Error
		if err != nil {
			return
		}

		// platform
		var p = struct {
			PlatformID  int      `json:"platform_id"`
			PlatformNum int      `json:"platform_num"`
			Platform    Platform `json:"platform"`
		}{}
		if userID != -1 {
			err = tx.Model(&Favorite{}).
				Where("user_id = ?", userID).
				Joins("left JOIN commodity_item ON commodity_item.id = favorite.commodity_item_id").
				Select("count(*) as platform_num, platform_id").
				Group("platform_id").
				Order("platform_num DESC").
				Limit(1).Scan(&p).Error
			// SELECT COUNT(*) AS platform_num, platform_id FROM favorite LEFT JOIN commodity_item ON commodity_item.id = favorite.commodity_item_id WHERE user_id = 2 GROUP BY platform_id ORDER BY platform_num DESC
		} else {
			err = tx.Model(&Favorite{}).
				Joins("left JOIN commodity_item ON commodity_item.id = favorite.commodity_item_id").
				Select("count(*) as platform_num, platform_id").
				Group("platform_id").
				Order("platform_num DESC").
				Limit(1).Scan(&p).Error
		}
		if err != nil {
			return
		}
		err = tx.Find(&p.Platform, p.PlatformID).Error
		if err != nil {
			return
		}

		// category
		var cat = struct {
			Category    string `json:"category"`
			CategoryNum int    `json:"category_num"`
		}{}
		if userID != -1 {
			err = tx.Model(&Favorite{}).
				Where("user_id = ?", userID).
				Joins("left JOIN commodity_item ON commodity_item.id = favorite.commodity_item_id").
				Joins("Left Join commodity ON commodity_item.commodity_id = commodity.id").
				Select("count(*) as category_num, category").
				Group("category").
				Order("category_num DESC").
				Limit(1).Scan(&cat).Error
			// SELECT COUNT(*) AS category_num, category FROM favorite LEFT JOIN commodity_item ON commodity_item.id = favorite.commodity_item_id LEFT JOIN commodity ON commodity_item.commodity_id = commodity.id WHERE user_id = 2 GROUP BY category ORDER BY category_num DESC
		} else {
			err = tx.Model(&Favorite{}).
				Joins("left JOIN commodity_item ON commodity_item.id = favorite.commodity_item_id").
				Joins("Left Join commodity ON commodity_item.commodity_id = commodity.id").
				Select("count(*) as category_num, category").
				Group("category").
				Order("category_num DESC").
				Limit(1).Scan(&cat).Error
		}
		if err != nil {
			return
		}

		res = SummaryResponse{
			Seller:         s.Seller,
			SellerNum:      s.SellerNum,
			Commodity:      c.Commodity,
			CommodityNum:   c.CommodityNum,
			Platform:       p.Platform,
			PlatformNum:    p.PlatformNum,
			Category:       cat.Category,
			CategoryNum:    cat.CategoryNum,
			FavoriteNum:    int(favoriteNum),
			MessageNum:     int(messageNum),
			PriceChangeNum: int(priceChangeNum),
		}
		return
	})
	return
}
