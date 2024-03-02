package json

import (
	"encoding/json"
	"fmt"
	"src/models"
)

func parsePlatform() error {
	data, err := ReadString(path + "platform.json")
	if err != nil {
		return err
	}
	var platforms []models.Platform
	err = json.Unmarshal([]byte(data), &platforms)
	if err != nil {
		return err
	}

	//err = models.DB.Create(&platforms).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	total := len(platforms)
	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}

		batch := platforms[i:end]
		err = models.DB.Create(&batch).Error
		if err != nil {
			fmt.Println(err)
			//return err
		}
	}

	return nil
}
