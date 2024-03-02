package json

import (
	"encoding/json"
	"fmt"
	"sort"
	"src/models"
)

func parsePriceChange() error {
	data, err := ReadString(path + "price_change.json")
	if err != nil {
		return err
	}
	var changes []models.PriceChange
	err = json.Unmarshal([]byte(data), &changes)
	if err != nil {
		return err
	}
	sort.Sort(models.ByUpdateAt(changes))
	//err = models.CreatePriceChanges(changes)
	//if err != nil {
	//	fmt.Println(err)
	//}
	total := len(changes)
	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}

		batch := changes[i:end]
		err = models.DB.Create(&batch).Error
		if err != nil {
			fmt.Println(err)
			//return err
		}
	}
	return nil
}
