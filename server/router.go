package server

import (
	"crud-categories/internal/handler"
	"crud-categories/internal/model"
	"crud-categories/internal/repository"
	"crud-categories/internal/service"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	router.GET("/categories", categoryHandler.GetAll)
	router.GET("/categories/:id", categoryHandler.GetByID)
	router.POST("/categories", categoryHandler.Create)
	router.PUT("/categories/:id", categoryHandler.Update)
	router.DELETE("/categories/:id", categoryHandler.Delete)

	router.GET("/health", func(c *gin.Context) {
		resp := model.Response{
			Message: "OK",
		}
		c.JSON(http.StatusOK, resp)
	})

	return router
}
