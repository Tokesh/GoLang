package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectGoLang/source/domain/entity"
	"projectGoLang/source/infrastructure/utils"
	"strconv"
)

func (c *Controller) AddDeliveryStatusController(ctx *gin.Context) {
	delivery := entity.Delivery{}
	err := ctx.BindJSON(&delivery)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Service.AddDeliveryStatus(delivery)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "created",
	})
}
func (c *Controller) DeleteDeliveryStatusController(ctx *gin.Context) {
	delivery := entity.Delivery{}
	err := ctx.BindJSON(&delivery)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Service.DeleteDeliveryStatus(delivery)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
	})
}

func (c *Controller) SelectDeliveryStatusController(ctx *gin.Context) {
	orderId := ctx.Param("id")
	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	deliveries, err := c.Service.SelectDelivery(orderIdInt)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"delivery": deliveries,
	})
}
func (c *Controller) SelectLastDeliveryStatusController(ctx *gin.Context) {
	orderId := ctx.Param("id")
	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	delivery, err := c.Service.SelectLastDelivery(orderIdInt)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"delivery": delivery,
	})
}
func (c *Controller) UpdateDeliveryStatusController(ctx *gin.Context) {
	delivery := entity.Delivery{}
	err := ctx.BindJSON(&delivery)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Service.UpdateDelivery(delivery)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "updated",
	})
}
