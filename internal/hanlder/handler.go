package hanlder

import (
	"github.com/gin-gonic/gin"
	"http_server/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateListById)
			lists.DELETE("/:id", h.deleteListById)
		}
	}
	return router
}
