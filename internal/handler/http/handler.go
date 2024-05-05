package http

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

func (h *Handler) ShowHomepage(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{})
}

func (h *Handler) GetOrderById(ctx *gin.Context) {
	orderId := ctx.Param("id")

	order, err := h.orderService.FindById(orderId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Order with id = " + orderId + " not found"})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/", h.ShowHomepage)

	api := router.Group("/api")
	{
		orders := api.Group("/order")
		{
			orders.GET("/:id", h.GetOrderById)
		}
	}

	router.LoadHTMLGlob("public/*.html")
	router.StaticFS("/static/js", http.Dir("./static/js"))
	router.StaticFS("/static/css", http.Dir("./static/css"))

	return router
}
