package json

import (
	"encoding/json"
	"fmt"
	"src/models"
	"src/utils"
)

func parseSeller() error {
	data, err := ReadString(path + "seller.json")
	if err != nil {
		return err
	}
	var sellers []models.Seller
	err = json.Unmarshal([]byte(data), &sellers)
	if err != nil {
		return err
	}
	for i := range sellers {
		sellers[i].Password = utils.MakePassword(sellers[i].Password)
	}
	//err = models.DB.Create(&sellers).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	total := len(sellers)
	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}

		batch := sellers[i:end]
		err = models.DB.Create(&batch).Error
		if err != nil {
			fmt.Println(err)
			//return err
		}
	}
	return nil
}
