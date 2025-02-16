package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"short-link/docs"
	"short-link/internal/core/config"
)

func setSwaggerRoutes(router *gin.RouterGroup, config config.Swagger) {
	if !config.Enable {
		return
	}

	docs.SwaggerInfo.Title = config.Info.Title
	docs.SwaggerInfo.Description = config.Info.Description
	docs.SwaggerInfo.Version = config.Info.Version
	docs.SwaggerInfo.Schemes = []string{config.Schemes}
	docs.SwaggerInfo.Host = config.Host

	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		config.Username: config.Password,
	}))

	authorized.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
