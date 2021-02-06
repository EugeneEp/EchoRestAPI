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

	go u.CreateUser()

	return utils.Response(c, map[string]interface{}{
		"success":true,
		"msg":"User has been created",
	}, nil)
}

// Получить юзеров
func GetUsers(c echo.Context) error {

	res, resErr := models.GetUsers()

	return utils.Response(c, res, resErr)
}

// Получить конкретного юзера
func GetUser(c echo.Context) error {

	var u models.User

	u.Id = c.Param("id")

	res, resErr := u.GetUser()

	return utils.Response(c, res, resErr)
}

// Обновить юзера
func UpdateUser(c echo.Context) error {

	var u models.User

	u.Id = c.Param("id")

	if err := c.Bind(&u); err != nil {
		return err
	}

	res, resErr := u.UpdateUser()

	return utils.Response(c, res, resErr)
}

// Удалить юзера
func DeleteUser(c echo.Context) error {
	var u models.User

	u.Id = c.Param("id")

	res, resErr := u.DeleteUser()

	return utils.Response(c, res, resErr)
}