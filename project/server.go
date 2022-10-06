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

	cfg                                         = entity.GetConfig()
	postgreSQLClient, err                       = postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	repository                                  = service.New(postgreSQLClient)
	Controller            controller.Controller = controller.New(repository)
)

func main() {
	server := gin.Default()
	//products
	server.GET("/products/:id", func(ctx *gin.Context) {
		ctx.JSON(200, Controller.FindOneProduct(ctx))
	})
	server.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, Controller.FindAllProducts(ctx))
	})
	server.POST("/products", func(ctx *gin.Context) {
		ctx.JSON(200, Controller.CreateOneProduct(ctx))
	})
	server.DELETE("/products/:id", func(ctx *gin.Context) {
		ctx.JSON(200, Controller.DeleteOneProduct(ctx))
	})
	//Login/Signup process
	server.POST("/signup", func(ctx *gin.Context) {
		ctx.JSON(200, Controller.SignUp(ctx))
	})
	server.POST("/login", func(ctx *gin.Context) {
		ctx.JSON(200, Controller.Login(ctx))
	})
	//Cart actions
	server.POST("/cart", func(ctx *gin.Context) {
		ctx.JSON(200, Controller.AddToCart(ctx))
	})
	server.DELETE("/cart", func(ctx *gin.Context) {
		ctx.JSON(200, Controller.DeleteFromCartByID(ctx))
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
