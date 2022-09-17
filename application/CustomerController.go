package application

import (
	"net/http"
	"ws_customers/domain/dto"
	"ws_customers/domain/service"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	orderService service.OrderService
}

func InitCustomerController(router *gin.Engine) {

	customerController := CustomerController{
		orderService: service.InitCustomertServiceImpl(),
	}

	router.POST("/order", customerController.CreateOrder)
	router.GET("/order/:idOrder", customerController.FindOrder)
	router.PUT("/order/cancel/:idOrder", customerController.CancelOrder) //Pendiente

}

func (r *CustomerController) CreateOrder(c *gin.Context) {

	var order dto.OrderDto

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.Response{})
		return
	}

	result, response := r.orderService.CreateOrders(order)

	if response.Status != http.StatusOK {
		c.JSON(response.Status, response)
		return
	}
	c.JSON(response.Status, result)
}

func (cd *CustomerController) FindOrder(c *gin.Context) {

	id := c.Param("idOrder")
	order := cd.orderService.FindOder(id)

	c.JSON(http.StatusOK, order)

}

func (cd *CustomerController) CancelOrder(c *gin.Context) {

	id := c.Param("idOrder")
	err, response := cd.orderService.CancelOrder(id)
	if response.Status != http.StatusOK {
		c.JSON(response.Status, err)
		return
	}
	c.JSON(http.StatusOK, response)

}
