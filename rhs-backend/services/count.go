package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	infra "rhs-backend/infraestructure"

	"github.com/gin-gonic/gin"
)

// CountServices returns a handler that reports the number of configured services.
//
//	@Summary		Count configured services
//	@Description	Returns the number of services stored at the path configured by SERVICE_PATH.
//	@Tags			services
//	@Produce		json
//	@Success		200	{integer}	integer
//	@Failure		500	"Service configuration could not be loaded"
//	@Router			/countServices [get]
func CountServices() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := infra.GetServiceFilePath()
		if path == "" {
			noPath := errors.New("Nao tem caminho pro arquivo de config.json")
			ctx.AbortWithError(http.StatusInternalServerError, noPath)
			return
		}

		data, err := os.ReadFile(path)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		var items []json.RawMessage

		if err := json.Unmarshal(data, &items); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		count := len(items)

		ctx.JSON(http.StatusOK, count)

	}
}
