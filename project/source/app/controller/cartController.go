package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectGoLang/source/infrastructure/utils"
	"strconv"
)

//	func (c *Controller) AddToCart(ctx *gin.Context) {
//		cart := entity.Cart{}
//		ctx.BindJSON(&cart)
//		return c.Service.AddToCart(&cart)
//	}
func (c *Controller) DeleteFromCartByID(ctx *gin.Context) {
	userId := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Service.DeleteCart(userIdInt)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
	})
}

func (c *Controller) SelectFromCartByID(ctx *gin.Context) {
	userId := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	carts, err := c.Service.SelectCart(userIdInt)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"cartOfUser": carts,
	})
}
