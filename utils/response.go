package utils

import (
	"net/http"
	"github.com/labstack/echo"
)

// Функция отправки ответа от сервера

func Response(c echo.Context, res interface{}, err error) error {
	if err != nil {
		return c.JSON(http.StatusBadRequest, res)
	}else{
		return c.JSON(http.StatusOK, res)
	}
}