package json

import (
	"encoding/json"
	"fmt"
	"src/models"
	"src/utils"
)

func parseUsers() error {
	data, err := ReadString(path + "user.json")
	if err != nil {
		return err
	}
	//fmt.Println("data:", data)

	//err = json.Unmarshal([]byte(data), &users)
	//fmt.Println(len(users))
	//for _, v := range users {
	//	var singleUser = models.User{
	//		ID:       0,
	//		Username: ,
	//		Password: "",
	//		Email:    "",
	//		Age:      0,
	//		Gender:   false,
	//		Phone:    "",
	//	}
	//}
	var users []models.User
	err = json.Unmarshal([]byte(data), &users)
	if err != nil {
		return err
	}
	for i := range users {
		users[i].Password = utils.MakePassword(users[i].Password)
		//fmt.Println(v)
		//return nil
	}

	total := len(users)
	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}

		batch := users[i:end]
		err = models.DB.Create(&batch).Error
		if err != nil {
			fmt.Println(err)
			//return err
		}
	}

	//err = models.DB.Create(&users).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	return nil
}
