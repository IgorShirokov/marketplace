package api

import (
	"github.com/IgorShirokov/marketplace/internal/catalog/api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	brands *handlers.BrandsHandler,
	categories *handlers.CategoriesHandler,
) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/brands", brands.GetBrands)
		v1.GET("/categories", categories.GetCategories)
	}
}
