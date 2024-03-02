package json

import (
	"encoding/json"
	"fmt"
	"sort"
	"src/models"
)

func parseMessage() error {
	data, err := ReadString(path + "message.json")
	if err != nil {
		return err
	}
	var message []models.Message
	err = json.Unmarshal([]byte(data), &message)
	if err != nil {
		return err
	}
	sort.Sort(models.ByCreatedAt(message))

	//err = models.DB.Create(&message).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	total := len(message)
	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}

		batch := message[i:end]
		err = models.DB.Create(&batch).Error
		if err != nil {
			fmt.Println(err)
			//return err
		}
	}
	return nil
}
