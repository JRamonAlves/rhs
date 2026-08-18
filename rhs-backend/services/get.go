package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	infra "rhs-backend/infraestructure"

	"github.com/gin-gonic/gin"
)

// Category identifies the group to which a service belongs.
type Category string

const (
	// MoviesAndSeries groups video streaming and library services.
	MoviesAndSeries Category = "Movies and Series"
	// Services groups general-purpose services.
	Services Category = "Services"
	// Photos groups photo management services.
	Photos Category = "Photos"
	// Comics groups comic and ebook services.
	Comics Category = "Comics"
)

// IsValid reports whether cat is one of the supported service categories.
func IsValid(cat Category) bool {
	switch cat {
	case MoviesAndSeries,
		Services,
		Photos,
		Comics:
		return true
	}
	return false

}

// Service describes a service exposed by the service configuration.
type Service struct {
	Name        string   `json:"name"`
	Url         string   `json:"url"`
	Port        int      `json:"port"`
	Description string   `json:"description"`
	Category    Category `json:"category"`
}

// GetServicesHandler returns the JSON service configuration from SERVICE_PATH.
//
//	@Summary		Get service configuration
//	@Description	Returns the JSON document stored at the path configured by SERVICE_PATH.
//	@Tags			services
//	@Produce		json
//	@Success		200	{object}	object
//	@Failure		500	"Service configuration could not be loaded"
//	@Router			/getServices [get]
func GetServicesHandler() gin.HandlerFunc {
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

		var services []Service
		if err := json.Unmarshal(data, &services); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		for _, service := range services {
			if !IsValid(service.Category) {
				err := errors.New("invalid service category: " + string(service.Category))
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}

		ctx.JSON(http.StatusOK, services)
	}
}
