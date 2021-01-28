package utils

import (
	"net/http"
	"github.com/labstack/echo"
)

// Функция отправки ответа от сервера

func Response(c echo.Context, res map[string]interface{}) error {
	if res["success"] == false {
		return c.JSON(http.StatusBadRequest, res)
	}else{
		return c.JSON(http.StatusOK, res)
	}
}