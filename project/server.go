package main

import (
	"context"
	"fmt"
	"projectGoLang/controller"
	"projectGoLang/entity"
	"projectGoLang/postgresql"
	"projectGoLang/service"

	"github.com/gin-gonic/gin"
)

var (
	productService        service.ProductService       = service.New()
	productController     controller.ProductController = controller.New(productService)
	cfg                                                = entity.GetConfig()
	postgreSQLClient, err                              = postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	repository                                         = postgresql.NewRepository(postgreSQLClient)
)

func main() {
	server := gin.Default()

	fmt.Println(repository.FindOne(context.TODO(), "2"))

	server.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, productController.FindAllProducts())
	})
	server.POST("/products", func(ctx *gin.Context) {
		fmt.Println(ctx)
		ctx.JSON(200, productController.AddProduct(ctx))
	})
	server.GET("/products/:id", func(ctx *gin.Context) {
		ctx.JSON(200, repository.FindOne(ctx, "5z`"))
	})
	server.Run(":8080")
}
