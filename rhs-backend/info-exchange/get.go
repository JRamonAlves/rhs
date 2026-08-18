package infoexchange

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// GetValuesHandler returns a handler that retrieves a value from Redis.
//
//	@Summary		Get a Redis value
//	@Description	Returns the value stored under the supplied Redis key.
//	@Tags			redis
//	@Produce		json
//	@Param			key	query		string	true	"Redis key"
//	@Success		200	{object}	map[string]string
//	@Failure		406	"Key is required"
//	@Router			/getValues [get]
func GetValuesHandler(db redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.Query("key")
		if key == "" {
			ctx.AbortWithStatus(http.StatusNotAcceptable)
			return
		}

		value := db.Get(ctx, key).Val()

		ctx.JSON(http.StatusOK, value)
	}
}
