package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"projectGoLang/source/app/controller"
	"projectGoLang/source/app/services"
	"projectGoLang/source/domain/entity"
	"projectGoLang/source/infrastructure/postgresql"
	"projectGoLang/source/infrastructure/repositories"
)

var (
	//productService        repositories.Repository           = repositories.New()

	cfg                   = entity.GetConfig()
	postgreSQLClient, err = postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	repository            = repositories.New(postgreSQLClient)
	service               = services.NewService(&repository)
	Controller            = controller.New(*service)
)

func main() {
	server := gin.Default()
	server.GET("/cart/:id", Controller.SelectFromCartByID)
	server.DELETE("/cart/:id", Controller.DeleteFromCartByID)
	server.Run()
	//if err := server.Run(":8080", Controller.InitRoutes()); err != nil {
	//	log.Fatalf("Error: %s", err.Error())
	//}
	//products
	//server.GET("/products/:id", func(ctx *gin.Context) {
	//	ctx.JSON(200, Controller.FindOneProduct(ctx))
	//})
	//server.GET("/products", func(ctx *gin.Context) {
	//	ctx.JSON(200, Controller.FindAllProducts(ctx))
	//})
	//server.POST("/products", func(ctx *gin.Context) {
	//	ctx.JSON(200, Controller.CreateOneProduct(ctx))
	//})
	//server.DELETE("/products/:id", func(ctx *gin.Context) {
	//	ctx.JSON(200, Controller.DeleteOneProduct(ctx))
	//})
	////Login/Signup process
	//server.POST("/signup", func(ctx *gin.Context) {
	//	ctx.JSON(200, Controller.SignUp(ctx))
	//})
	//server.POST("/login", func(ctx *gin.Context) {
	//	ctx.JSON(200, Controller.Login(ctx))
	//})
	////Cart actions
	//server.POST("/cart", func(ctx *gin.Context) {
	//	ctx.JSON(200, Controller.AddToCart(ctx))
	//})
	//server.DELETE("/cart", func(ctx *gin.Context) {
	//	ctx.JSON(200, Controller.DeleteFromCartByID(ctx))
	//})
	//server.GET("/cart", func(ctx *gin.Context) {
	//	ctx.
	//})
	//server.POST("/products", func(ctx *gin.Context) {
	//	fmt.Println(ctx)
	//	ctx.JSON(200, productController.AddProduct(ctx))
	//})
	//server.GET("/products/:id", func(ctx *gin.Context) {
	//	ctx.JSON(200, repository.FindOne(ctx))
	//})

}
