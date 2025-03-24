package main

import (
	_ "sparrow/docs"
	"sparrow/handlers"
	"sparrow/utils"

	"github.com/gin-contrib/cors"
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
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	spifs := utils.LoadPolicies()
	// Swagger endpoint
	r.GET("/documentation/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Playground route
	r.Static("/static", "./playground/build/static")
	r.GET("/playground/*path", func(c *gin.Context) {
		c.File("./playground/build/index.html")
	})

	// GET API Routes
	// TODO use r.Groups
	r.GET("/api/v1/policies", handlers.PoliciesHandler(spifs))
	r.GET("/api/v1/classifications/:policy", handlers.ClassificationsHandler(spifs))
	r.GET("/api/v1/categories/:policy/*classification", handlers.CategoriesHandler(spifs))
	r.GET("/api/v1/type/:policy/:category", handlers.TypeHandler(spifs))
	r.GET("/api/v1/mentions/:policy/:classification/:category", handlers.MentionsHandler(spifs))

	// POST API routes
	//r.POST("/api/v1/marking/:type", handlers.MarkingHandler())
	//r.POST("/api/v1/dominant/", handlers.DominantLabelHandler())
	r.POST("/api/v1/generate", handlers.GenerateHandler())
	r.POST("/api/v1/parse", handlers.ParseHandler())

	r.Run(":8080") //r.RunTLS(crt,key)
}
