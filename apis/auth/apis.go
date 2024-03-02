package auth

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	"gorm.io/gorm"
	. "src/models"
	"src/utils"
	"strings"
)

// Login godoc
// @Router /api/login [post]
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param json body LoginRequest true "json"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} common.HttpError
// @Failure 401 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @Failure 500 {object} common.HttpError
func Login(c *fiber.Ctx) error {
	var body LoginRequest
	err := common.ValidateBody(c, &body)
	if err != nil {
		return err
	}

	var tmpUser TmpUser
	var result *gorm.DB
	if strings.ToLower(body.Type) == "seller" {
		var seller Seller
		result = DB.Where("username = ?", body.Username).First(&seller)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return common.Unauthorized("账号未注册")
			}
			return common.InternalServerError()
		}
		tmpUser.ID = seller.ID
		tmpUser.Username = seller.Username
		tmpUser.UserType = "seller"
		tmpUser.Password = seller.Password
	} else if strings.ToLower(body.Type) == "admin" {
		var admin Admin
		result = DB.Where("username = ?", body.Username).First(&admin)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return common.Forbidden("无权限")
			}
			return common.InternalServerError()
		}
		tmpUser.ID = admin.ID
		tmpUser.Username = admin.Username
		tmpUser.UserType = "admin"
		tmpUser.Password = admin.Password
	} else {
		var user User
		result = DB.Where("username = ?", body.Username).First(&user)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return common.Unauthorized("账号未注册")
			}
			return common.InternalServerError()
		}
		tmpUser.ID = user.ID
		tmpUser.Username = user.Username
		tmpUser.UserType = "user"
		tmpUser.Password = user.Password
	}

	ok := utils.CheckPassword(body.Password, tmpUser.Password)
	if !ok {
		//return fiber.NewError(fiber.StatusUnauthorized, "密码错误")
		return common.Unauthorized("密码错误")
	}

	access, err := tmpUser.CreateJWTToken()
	if err != nil {
		return err
	}

	//// update login time
	//err = DB.Model(&user).Select("LastLogin").Updates(&user).Error
	//if err != nil {
	//	return err
	//}

	return c.JSON(TokenResponse{
		Access:  access,
		Message: "登录成功",
	})
}

func Logout(c *fiber.Ctx) error {
	return nil
}

//为什么以下不在user里面
//Register
//ChangePassword
//ChangeEmail
