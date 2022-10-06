package controller

import (
	"github.com/gin-gonic/gin"
	"projectGoLang/entity"
)

func (c *Controller) AddToCart(ctx *gin.Context) *entity.Cart {
	cart := entity.Cart{}
	ctx.BindJSON(&cart)
	return c.Repository.AddToCart(&cart)
}

func (c *Controller) DeleteFromCartByID(ctx *gin.Context) *entity.Cart {
	user_id := entity.Cart{}
	ctx.BindJSON(&user_id)
	return c.Repository.DeleteFromCartByUserID(&user_id)
}
