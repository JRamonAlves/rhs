package infoexchange

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// SetValuesHandler returns a handler that stores a value in Redis.
//
//	@Summary		Set a Redis value
//	@Description	Stores a value under the supplied Redis key without expiration.
//	@Tags			redis
//	@Param			key		query	string	true	"Redis key"
//	@Param			value	query	string	true	"Value to store"
//	@Success		200		"Value stored"
//	@Failure		406		"Key and value are required"
//	@Failure		500		"Redis operation failed"
//	@Router			/setValues [post]
func SetValuesHandler(db redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.Query("key")
		value := ctx.Query("value")
		if key == "" || value == "" {
			ctx.AbortWithStatus(http.StatusNotAcceptable)
			return
		}

		err := db.Set(ctx, key, value, 0).Err()
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		ctx.Status(200)
	}
}
