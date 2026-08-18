package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	goredis "github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "rhs-backend/docs"
	infoexchange "rhs-backend/info-exchange"
	infra "rhs-backend/infraestructure"
	"rhs-backend/services"

	"github.com/gin-gonic/gin"
)

//	@title			RHS Backend API
//	@version		1.0
//	@description	This is the RHS Backend API documentation.
//
// main configures the HTTP routes and starts the API server.
func main() {
	_ = godotenv.Load()

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/ping", pingHandler)

	// Carrega o swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api/v1")
	{
		v1.GET("/hello", HelloHandler)
	}

	// Redis-like service for strings
	redisConfig := infra.RedisConfig()
	redisClient := goredis.NewClient(&redisConfig)
	router.POST("/setValues", infoexchange.SetValuesHandler(*redisClient))
	router.GET("/getValues", infoexchange.GetValuesHandler(*redisClient))

	// Get the data of the services running
	router.GET("/getServices", services.GetServicesHandler())
	router.GET("/countServices", services.CountServices())

	router.Run() // listens on 0.0.0.0:8080 by default
}

// pingHandler reports whether the API is running.
//
//	@Summary		Check API health
//	@Description	Returns a pong response when the API is running.
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/ping [get]
func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// HelloHandler returns a greeting message.
//
//	@Summary		Show a hello message
//	@Description	Returns a greeting from the server.
//	@Tags			example
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/api/v1/hello [get]
func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}
