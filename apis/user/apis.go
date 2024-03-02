package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opentreehole/go-common"
	. "src/models"
	"src/utils"
)

//// GetCurrentUser @GetCurrentUser
//func GetCurrentUser(c *fiber.Ctx) error {
//	user, err := GetGeneralUser(c)
//	if err != nil {
//		return err
//	}
//	return c.JSON(&user)
//}

// GetUserInfo @GetUserInfo
// @Router /api/users/data [get]
// @Summary Get user info
// @Description Get user info
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func GetUserInfo(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}

	if tmpUser.UserType != "user" {
		return common.Forbidden()
	}

	getUser, err := GetUserByID(tmpUser.ID)
	if err != nil {
		return err
	}

	return c.JSON(&getUser)
}

// GetAllUsers @GetAllUsers
// @Router /api/users/all [get]
// @Summary Get all users
// @Description Get all users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func GetAllUsers(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" {
		return common.Forbidden("Only admin can get all users")
	}
	users, err := GetUsers()
	if err != nil {
		return err
	}
	return c.JSON(&users)
}

// AddUser @AddUser
// @Router /api/users [post]
// @Summary Add user
// @Description Add user
// @Tags User
// @Accept json
// @Produce json
// @Param json body CreateUserRequest true "json"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func AddUser(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" {
		return common.Forbidden("Only admin can add user")
	}
	var user User
	err = common.ValidateBody(c, &user)
	if err != nil {
		return err
	}
	user.Password = utils.MakePassword(user.Password)
	return user.Create()
}

// UpdateUser @UpdateUser
// @Router /api/users [put]
// @Summary Update user
// @Description Update user
// @Tags User
// @Accept json
// @Produce json
// @Param json body models.User true "json"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @Failure 404 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func UpdateUser(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" && tmpUser.UserType != "user" {
		return common.Forbidden("Only admin or user can update user")
	}
	var user User
	err = common.ValidateBody(c, &user)
	if err != nil {
		return err
	}
	if tmpUser.UserType == "user" && tmpUser.ID != user.ID {
		return common.Forbidden("You can only update your own info")
	}
	if len(user.Password) != 0 {
		user.Password = utils.MakePassword(user.Password)
	}
	return user.Update()
}

// DeleteUser @DeleteUser
// @Router /api/users [delete]
// @Summary Delete user
// @Description Delete user
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Success 200
// @Failure 400 {object} common.HttpError
// @Failure 403 {object} common.HttpError
// @Failure 404 {object} common.HttpError
// @param Authorization header string true "Bearer和token空格拼接"
func DeleteUser(c *fiber.Ctx) error {
	tmpUser, err := GetGeneralUser(c)
	if err != nil {
		return err
	}
	if tmpUser.UserType != "admin" {
		return common.Forbidden("Only admin can delete user")
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	return DeleteUserByID(id)
}
