package routes

import (
	"github.com/MaksimUlitin/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.singIn)
		auth.POST("/sign-in", h.singIn)

	}
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllList)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deletList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createitem)
				items.GET("/", h.getAllitem)
				items.GET("/:item_id", h.getAllitem)
				items.PUT("/:item_id", h.getitemById)
				items.DELETE("/:item_id", h.deletitem)

			}
		}
	}
	return router
}
