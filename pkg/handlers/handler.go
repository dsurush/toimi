package handlers

import (
	"advert/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		advert := api.Group("/advert")
		{
			advert.POST("/create", h.create)
			advert.GET("/:id", h.getById)
			advert.GET("/all", h.getAll)
		}
	}

	return router
}
