package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectGoLang/source/domain/entity"
	"projectGoLang/source/infrastructure/utils"
	"strconv"
)

func (c *Controller) AddToPayment(ctx *gin.Context) {
	payment := entity.Payment{}
	err := ctx.BindJSON(&payment)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Service.AddToPayment(payment)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "created",
	})
}
func (c *Controller) DeleteFromPayment(ctx *gin.Context) {
	paymentId := ctx.Param("id")
	paymentIdInt, err := strconv.Atoi(paymentId)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Service.DeletePayment(paymentIdInt)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
	})
}

func (c *Controller) SelectPaymentByPaymentID(ctx *gin.Context) {
	paymentId := ctx.Param("id")
	paymentIdInt, err := strconv.Atoi(paymentId)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	payment, err := c.Service.SelectPaymentByPaymentID(paymentIdInt)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Payment": payment,
	})
}
func (c *Controller) SelectPaymentByUserID(ctx *gin.Context) {
	userId := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	payments, err := c.Service.SelectPaymentByUserID(userIdInt)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Payment": payments,
	})
}
func (c *Controller) UpdatePaymentControl(ctx *gin.Context) {
	payment := entity.Payment{}
	err := ctx.BindJSON(&payment)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Service.UpdatePayment(payment)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "updated",
	})
}
