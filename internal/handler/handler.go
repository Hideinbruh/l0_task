package handler

import (
	"awesomeProject/cache"
	"awesomeProject/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	cache    *cache.Cache
}

func NewHanlder(services *service.Service, cache *cache.Cache) *Handler {
	return &Handler{services: services, cache: cache}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	order := router.Group("/order")
	{
		order.POST("/", h.Get)
		order.GET("/", h.CreateHTML)
	}
	return router
}
