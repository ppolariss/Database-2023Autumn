package json

import (
	"encoding/json"
	"fmt"
	"src/models"
	"src/utils"
)

func parseAdmin() error {
	data, err := ReadString(path + "admin.json")
	if err != nil {
		return err
	}
	var admins []models.Admin
	err = json.Unmarshal([]byte(data), &admins)
	if err != nil {
		return err
	}
	for i := range admins {
		admins[i].Password = utils.MakePassword(admins[i].Password)
	}
	//err = models.DB.Create(&admins).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	total := len(admins)
	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}

		batch := admins[i:end]
		err = models.DB.Create(&batch).Error
		if err != nil {
			fmt.Println(err)
			//return err
		}
	}
	return nil
}
