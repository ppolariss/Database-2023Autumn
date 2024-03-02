package json

import (
	"encoding/json"
	"fmt"
	"src/models"
)

func parseFavorite() error {
	data, err := ReadString(path + "favorite.json")
	if err != nil {
		return err
	}
	var favorite []models.Favorite
	err = json.Unmarshal([]byte(data), &favorite)
	if err != nil {
		return err
	}

	//err = models.DB.Create(&favorite).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	total := len(favorite)
	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}

		batch := favorite[i:end]
		err = models.DB.Create(&batch).Error
		if err != nil {
			fmt.Println(err)
			//return err
		}
	}

	return nil
}
