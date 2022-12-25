package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectGoLang/source/domain/entity"
	"projectGoLang/source/infrastructure/utils"
)

func (c *Controller) AddStoreInventoryController(ctx *gin.Context) {
	storeInventory := entity.StoreInventory{}
	err := ctx.BindJSON(&storeInventory)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Service.AddStoreInventoryService(storeInventory)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "created",
	})
}
func (c *Controller) DeleteStoreInventoryController(ctx *gin.Context) {
	storeInventory := entity.StoreInventory{}
	err := ctx.BindJSON(&storeInventory)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Service.DeleteStoreInventoryService(storeInventory.StoreId, storeInventory.ProductId)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
	})
}

func (c *Controller) SelectStoreInventoryController(ctx *gin.Context) {
	storeInventory := entity.StoreInventory{}
	err := ctx.BindJSON(&storeInventory)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	storeInventor, err := c.Service.SelectStoreInventoryService(storeInventory.StoreId, storeInventory.ProductId)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"delivery": storeInventor,
	})
}
func (c *Controller) UpdateStoreInventoryController(ctx *gin.Context) {
	storeInventory := entity.StoreInventory{}
	err := ctx.BindJSON(&storeInventory)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Service.UpdateStoreInventoryService(storeInventory)
	if err != nil {
		utils.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "updated",
	})
}
