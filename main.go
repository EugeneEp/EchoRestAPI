package main

import (
	"ECHORestAPI/router"
	"ECHORestAPI/models"
)

func main() {
	// Инициализируем роутер
	e := router.InitRouter()
	models.InitDB() // Подключаем базу данных (json файл)
	e.Logger.Fatal(e.Start(":1323")) // Запускаем сервер
}