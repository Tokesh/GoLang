package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"projectGoLang/controller"
	"projectGoLang/entity"
	"projectGoLang/postgresql"
	"projectGoLang/service"
)

var (
	//productService        service.Repository           = service.New()

	cfg                                                = entity.GetConfig()
	postgreSQLClient, err                              = postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	repository                                         = service.New(postgreSQLClient)
	productController     controller.ProductController = controller.New(repository)
)

func main() {
	server := gin.Default()

	server.GET("/products/:id", func(ctx *gin.Context) {
		ctx.JSON(200, productController.FindOneProduct(ctx))
	})
	server.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, productController.FindAllProducts(ctx))
	})
	server.POST("/products", func(ctx *gin.Context) {
		ctx.JSON(200, productController.CreateOneProduct(ctx))
	})
	server.DELETE("/products/:id", func(ctx *gin.Context) {
		ctx.JSON(200, productController.DeleteOneProduct(ctx))
	})
	//server.POST("/products", func(ctx *gin.Context) {
	//	fmt.Println(ctx)
	//	ctx.JSON(200, productController.AddProduct(ctx))
	//})
	//server.GET("/products/:id", func(ctx *gin.Context) {
	//	ctx.JSON(200, repository.FindOne(ctx))
	//})
	server.Run(":8080")
}
