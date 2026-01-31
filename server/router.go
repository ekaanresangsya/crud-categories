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

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo, categoryRepo)
	productHandler := handler.NewProductHandler(productService)

	categories := router.Group("/categories")
	{
		categories.GET("/", categoryHandler.GetAll)
		categories.GET("/:id", categoryHandler.GetByID)
		categories.POST("/", categoryHandler.Create)
		categories.PUT("/:id", categoryHandler.Update)
		categories.DELETE("/:id", categoryHandler.Delete)
	}

	products := router.Group("/products")
	{
		products.GET("/", productHandler.GetAll)
		products.GET("/:id", productHandler.GetByID)
		products.POST("/", productHandler.Create)
		products.PUT("/:id", productHandler.Update)
		products.DELETE("/:id", productHandler.Delete)
	}

	router.GET("/health", func(c *gin.Context) {
		resp := model.Response{
			Message: "OK",
		}
		c.JSON(http.StatusOK, resp)
	})

	return router
}
