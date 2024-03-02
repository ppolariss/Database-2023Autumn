package json

import (
	"encoding/json"
	"fmt"
	"src/models"
)

func parseItem() error {
	data, err := ReadString(path + "commodity_item.json")
	if err != nil {
		return err
	}
	var items []models.CommodityItem
	err = json.Unmarshal([]byte(data), &items)
	if err != nil {
		return err
	}
	//err = models.CreateItems(items)
	//if err != nil {
	//	fmt.Println(err)
	//}
	total := len(items)
	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}

		batch := items[i:end]
		err = models.DB.Create(&batch).Error
		if err != nil {
			fmt.Println(err)
			//return err
		}
	}
	return nil
}
