package main

import (
	_ "sparrow/docs"
	"sparrow/handlers"
	"sparrow/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title SPARROW API Documentation
// @version 1.0
// @description SPARROW Project API Documentation generated using Swagger
// @BasePath /
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	spifs := utils.LoadPolicies()
	// Swagger endpoint
	r.GET("/documentation/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// GET API Routes
	r.GET("/api/v1/policies", handlers.PoliciesHandler(spifs))
	r.GET("/api/v1/classifications/:policy", handlers.ClassificationsHandler(spifs))
	r.GET("/api/v1/categories/:policy/*classification", handlers.CategoriesHandler(spifs))
	r.GET("/api/v1/type/:policy/:category", handlers.TypeHandler(spifs))
	r.GET("/api/v1/mentions/:policy/:classification/:category", handlers.MentionsHandler(spifs))

	// POST API routes
	//r.POST("/api/v1/marking/:type", handlers.MarkingHandler())
	//r.POST("/api/v1/dominant/", handlers.DominantLabelHandler())
	//r.POST("/api/v1/generate", handlers.GenerateHandler())
	r.POST("/api/v1/parse", handlers.ParseHandler())

	r.Run(":8080") //r.RunTLS(crt,key)
}
