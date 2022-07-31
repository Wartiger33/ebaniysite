package handler

import (
	"awesomeProject1/pkg/service"
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
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllList)
			lists.GET("/:id", h.getListsById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deliteList)

			items := lists.Group("/")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:items_id", h.getItemsById)
				items.PUT("/:items_id", h.updateList)
				items.DELETE("/:items_id", h.deliteItem)
			}
		}
	}
	return router
}
