package main

import (
	"minibook-backend/internal/db" //Пакет, содержащий логику для работы с базой данных
	// (например, инициализация подключения к PostgreSQL и создание таблиц
	"minibook-backend/internal/handlers" //Пакет с обработчиками HTTP-запросов
	// (например, для регистрации, входа и управления транзакциями

	"github.com/gin-gonic/gin" //Фреймворк Gin для создания веб-сервера и маршрутизации
)

func main() {
	// Инициализация БД
	db.InitDB() //Вызывает функцию InitDB из пакета internal/db. Эта функция,
	// вероятно, устанавливает соединение с базой данных PostgreSQL и выполняет
	// инициализацию (например, создаёт таблицы, как указано в createTables в предыдущих ошибках)

	// Настройка роутера
	r := gin.Default() //Создаёт экземпляр роутера Gin с middleware Logger (для логирования запросов)
	// и Recovery (для обработки паник).

	// Группа API
	api := r.Group("/api/v1") //Создаёт группу маршрутов с префиксом /api/v1.
	{
		// Аутентификация
		api.POST("/auth/register", handlers.Register) //Регистрирует нового пользователя.
		api.GET("/auth/login", handlers.Login)        //Аутентифицирует пользователя и возвращает JWT-токен

		// Защищенные маршруты (нужен middleware). Создаёт подгруппу маршрутов, защищённых middleware AuthMiddleware,
		// который проверяет JWT-токен.
		protected := api.Group("/")
		protected.Use(handlers.AuthMiddleware())
		{
			protected.POST("/transactions", handlers.CreateTransaction) //Создаёт новую транзакцию.
			protected.GET("/transactions", handlers.GetTransactions)    //Возвращает список транзакций пользователя.
		}
	}

	// Запуск сервера
	r.Run(":8080") //Запускает сервер на порту 8080 (http://localhost:8080).
}
