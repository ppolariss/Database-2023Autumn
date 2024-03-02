package json

import (
	"encoding/json"
	"fmt"
	"src/models"
)

func parseCommodity() error {
	data, err := ReadString(path + "commodity.json")
	if err != nil {
		return err
	}
	var commodities []models.Commodity
	err = json.Unmarshal([]byte(data), &commodities)
	if err != nil {
		return err
	}

	//err = models.DB.Create(&commodities).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	total := len(commodities)
	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}

		batch := commodities[i:end]
		err = models.DB.Create(&batch).Error
		if err != nil {
			fmt.Println(err)
			//return err
		}
	}

	return nil
}
