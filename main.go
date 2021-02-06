package main

import (
	"ECHORestAPI/router"
	"ECHORestAPI/models"
)

func main() {
	// Инициализируем роутер
	e := router.InitRouter()
	models.InitDB() // Подключаем базу данных (json файл)
	go models.Db.RunDBHub() // Прослушиваем каналы для работы с бд
	e.Logger.Fatal(e.Start(":1323")) // Запускаем сервер
}