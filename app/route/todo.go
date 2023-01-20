package route

import (
	"github.com/labstack/echo/v4"
	"github.com/vandyahmad24/moonlay-test/app/handler"
	"github.com/vandyahmad24/moonlay-test/app/middleware"
	"github.com/vandyahmad24/moonlay-test/app/repository"
	"github.com/vandyahmad24/moonlay-test/app/usecase"
	"github.com/vandyahmad24/moonlay-test/config/database"
)

func RouteTodo(e *echo.Echo) {
	repo := repository.NewRepository(database.DB)
	uc := usecase.NewUsecase(repo)
	route := handler.NewTodoHanler(uc)

	e.GET("/todo", route.GetAllParent)
	e.GET("/todo-with-child", route.GetAllWithChild)
	e.GET("/todo/:id", route.GetDetailTodo)
	e.POST("/todo", route.CreateTodo, middleware.FileTypeValidator([]string{".txt", ".pdf"}))
	e.POST("/sublist/:id", route.CreateSublist, middleware.FileTypeValidator([]string{".txt", ".pdf"}))
	e.PUT("/todo/:id", route.Update, middleware.FileTypeValidator([]string{".txt", ".pdf"}))
	e.DELETE("/todo/:id", route.Delete)
}
