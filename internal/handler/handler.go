package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orders_service/internal/service"
)

type Handler struct {
	orderService *service.OrderService
}

func NewHandler(s *service.OrderService) *Handler {
	return &Handler{orderService: s}
}

func (h *Handler) GetOrderById(ctx *gin.Context) {
	orderId := ctx.Param("id")

	order, err := h.orderService.FindById(orderId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order with id = " + orderId + " not found"})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		orders := api.Group("/order")
		{
			orders.GET("/:id", h.GetOrderById)
		}
	}

	return router
}
