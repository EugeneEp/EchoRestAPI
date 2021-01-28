package router

import(
	"ECHORestAPI/controllers"
	"github.com/labstack/echo"
)

// Инициализация роутера

func InitRouter()*echo.Echo{
	e := echo.New()

	e.POST("/users/create", controllers.CreateUser) // Создать юзера
	e.GET("/users/", controllers.GetUsers) // Получить всех юзеров
	e.GET("/users/:id", controllers.GetUser) // Получить юзера по id
	e.PUT("/users/:id/update", controllers.UpdateUser) // Обновить юзера
	e.DELETE("/users/:id/delete", controllers.DeleteUser) // Удалить юзера

	return e
}