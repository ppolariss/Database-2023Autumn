package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	"gorm.io/gorm"
)

type TmpUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

// GetGeneralUser
// return user from fiber.Ctx or jwt
func GetGeneralUser(c *fiber.Ctx) (user *TmpUser, err error) {
	//fmt.Println(c.Get("type"))

	if c.Locals("user") != nil {
		user = c.Locals("user").(*TmpUser)
		return
	}

	// get id and user_type from jwt
	token := common.GetJWTToken(c)
	if token == "" {
		return nil, common.Unauthorized("Unauthorized")
	}
	err = common.ParseJWTToken(token, &user)
	if err != nil {
		return nil, common.Unauthorized("Unauthorized")
	}

	// load user from database in transaction
	err = user.CheckUserID()

	//if user.IsAdmin {
	//	user.Permission.Admin = maxTime
	//} else {
	//	user.Permission.Admin = minTime
	//}
	//user.Permission.Silent = user.BanDivision
	//user.Permission.OffenseCount = user.OffenceCount
	//
	//if config.Config.UserAllShowHidden {
	//	user.Config.ShowFolded = "hide"
	//}

	if err != nil {
		return
	}
	// save user in c.Locals
	c.Locals("user", user)
	return
}

func (user *TmpUser) CheckUserID() error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if user.UserType == "seller" {
			var seller = Seller{ID: user.ID}
			err := tx.Take(&seller).Error
			if err != nil {
				return err
			}
			user.Username = seller.Username
		} else if user.UserType == "admin" {
			var admin = Admin{ID: user.ID}
			err := tx.Take(&admin).Error
			if err != nil {
				return err
			}
			user.Username = admin.Username
		} else if user.UserType == "user" {
			var newUser = User{ID: user.ID}
			err := tx.Take(&newUser).Error
			if err != nil {
				return err
			}
			user.Username = newUser.Username
		} else {
			return common.InternalServerError("未知用户类型")
		}
		//err := tx.Take(&user, userID).Error
		//if err != nil {
		//	if errors.Is(err, gorm.ErrRecordNotFound) {
		//		// insert user if not found
		//		user.ID = userID
		//		user.Config = defaultUserConfig
		//		err = tx.Create(&user).Error
		//		if err != nil {
		//			return err
		//		}
		//	} else {
		//		return err
		//	}
		//}

		return nil
	})
}
