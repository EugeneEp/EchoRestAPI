package controllers

import(
	"ECHORestAPI/models"
	"ECHORestAPI/utils"
	"github.com/labstack/echo"
)

// Создать юзера
func CreateUser(c echo.Context) error {
	var u models.User

	if err := c.Bind(&u); err != nil {
		return err
	}

	res := u.CreateUser()

	return utils.Response(c, res)
}

// Получить юзеров
func GetUsers(c echo.Context) error {

	res := models.GetUsers()

	return utils.Response(c, res)
}

// Получить конкретного юзера
func GetUser(c echo.Context) error {

	var u models.User

	u.Id = c.Param("id")

	res := u.GetUser()

	return utils.Response(c, res)
}

// Обновить юзера
func UpdateUser(c echo.Context) error {

	var u models.User

	u.Id = c.Param("id")

	if err := c.Bind(&u); err != nil {
		return err
	}

	res := u.UpdateUser()

	return utils.Response(c, res)
}

// Удалить юзера
func DeleteUser(c echo.Context) error {
	var u models.User

	u.Id = c.Param("id")

	res := u.DeleteUser()

	return utils.Response(c, res)
}