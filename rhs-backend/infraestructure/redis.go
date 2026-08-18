package infraestructure

import (
	"os"

	"github.com/gin-gonic/gin"
	goredis "github.com/redis/go-redis/v9"
)

// Config returns the Redis options for the current Gin mode.
func RedisConfig() goredis.Options {
	var addr string
	if gin.Mode() == gin.DebugMode {
		addr = "localhost:6379"
	} else {
		addr = "redis:6379"
	}

	password := os.Getenv("REDIS_PASSWORD")

	return goredis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	}
}
