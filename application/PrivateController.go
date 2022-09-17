package application

import (
	"ws_customers/domain/service"

	"github.com/gin-gonic/gin"
)

type PrivateController struct {
	privateService service.PrivateService
}

func InitPrivateController(router *gin.Engine) {

	privateController := PrivateController{
		privateService: service.InitPrivateServiceImpl(),
	}

	router.PUT("/order/status/:idOrder/:status", privateController.ChangeStatus) //Pendiente
}

func (r *PrivateController) ChangeStatus(c *gin.Context) {

	id := c.Param("idOrder")
	status := c.Param("status")

	response := r.privateService.ChangeStatus(id, status)

	c.JSON(response.Status, response)
}
