package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectGoLang/source/domain/entity"
	"projectGoLang/source/infrastructure/utils"
	"strconv"
)

func (c *Controller) AddOrderController(ctx *gin.Context) {
	order := entity.Order{}
	err := ctx.BindJSON(&order)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Service.AddToOrder(order)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	//cart, err := c.Service.SelectCart(order.ClientId)
	//err = c.Service.FillOrderService(cart, order.ClientId)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "created",
	})
}
func (c *Controller) DeleteOrderController(ctx *gin.Context) {
	orderId := ctx.Param("id")
	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Service.DeleteOrder(orderIdInt)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
	})
}

func (c *Controller) SelectOrderController(ctx *gin.Context) {
	orderId := ctx.Param("id")
	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	order, err := c.Service.SelectOrder(orderIdInt)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"delivery": order,
	})
}
func (c *Controller) UpdateOrderController(ctx *gin.Context) {
	order := entity.Order{}
	err := ctx.BindJSON(&order)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Service.UpdateOrder(order)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "updated",
	})
}
