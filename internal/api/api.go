package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"simple-service/internal/api/middleware"
	"simple-service/internal/service"
	"strconv"
	"simple-service/internal/dto"
)

// Routers - структура для хранения зависимостей роутов
type Routers struct {
	Service service.Service
}

func (r *Routers) GetTaskHandler(c *fiber.Ctx) error {
	// Получение ID из URL параметра
	idStr := c.Params("id")

	// Преобразование строки в число
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return dto.BadResponseError(c, dto.FieldIncorrect, "Invalid task ID")
	}

	// Вызов метода сервиса
	return r.Service.GetTask(c, id)
}


// NewRouters - конструктор для настройки API
func NewRouters(r *Routers, token string) *fiber.App {
	app := fiber.New()

	// Настройка CORS (разрешенные методы, заголовки, авторизация)
	app.Use(cors.New(cors.Config{
		AllowMethods:  "GET, POST, PUT, DELETE",
		AllowHeaders:  "Accept, Authorization, Content-Type, X-CSRF-Token, X-REQUEST-ID",
		ExposeHeaders: "Link",
		MaxAge:        300,
	}))

	// Группа маршрутов с авторизацией
	apiGroup := app.Group("/v1", middleware.Authorization(token))

	// Роут для создания задачи
	apiGroup.Post("/create_task", r.Service.CreateTask)

	apiGroup.Get("/task/:id", r.GetTaskHandler)


	return app
}
