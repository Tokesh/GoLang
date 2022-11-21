package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectGoLang/app/controller"
	"projectGoLang/source/app/services"
	"projectGoLang/source/domain/entity"
	"projectGoLang/source/infrastructure/postgresql"
	"projectGoLang/source/infrastructure/repositories"
	"projectGoLang/source/infrastructure/utils"
)

var (
	//productService repositories.Repository = repositories.New()

	cfg                   = entity.GetConfig()
	postgreSQLClient, err = postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	repository            = repositories.New(postgreSQLClient)
	service               = services.NewService(&repository)
	redisCache            = utils.NewRedisCache("localhost:6379", 0, 3600)
	Controller            = controller.New(*service, *redisCache)
)

func main() {
	server := gin.Default()
	//server.GET("/cart/:id", Controller.SelectFromCartByID)
	//server.DELETE("/cart/:id", Controller.DeleteFromCartByID)
	//server.POST("/cart", Controller.AddToCart)
	//server.PUT("/cart", Controller.UpdateCartControl)
	//http.HandleFunc("", homePage)
	// "RRR"
	cart := server.Group("/cart")
	{
		cart.GET("/:id", Controller.SelectFromCartByID)
		cart.DELETE("/:id", Controller.DeleteFromCartByID)
		cart.POST("", Controller.AddToCart)
		cart.PUT("", Controller.UpdateCartControl)
	}
	products := server.Group("/products")
	{
		products.GET("/:id", Controller.FindOneProduct)
		products.GET("", Controller.FindAllProducts)
		products.POST("", Controller.CreateOneProduct)
		products.PUT("", Controller.UpdateOneProduct)
		products.DELETE("/:id", Controller.DeleteOneProduct)
	}
	payments := server.Group("/payments")
	{
		payments.GET("/paymentId/:id", Controller.SelectPaymentByPaymentID)
		payments.GET("/userId/:id", Controller.SelectPaymentByUserID)
		payments.POST("", Controller.AddToPayment)
		payments.PUT("", Controller.UpdatePaymentControl)
		payments.DELETE("/:id", Controller.DeleteFromPayment)
	}
	server.POST("/user", Controller.Login)
	server.GET("/user", Controller.Validate)
	//"RRR"
	//server.POST("/login", Controller.Login)
	//server.GET("/products/:id", Controller.FindOneProduct)
	//server.GET("/products", Controller.FindAllProducts)
	//server.POST("/products", Controller.CreateOneProduct)
	//server.PUT("/products", Controller.UpdateOneProduct)
	//server.DELETE("/products/:id", Controller.DeleteOneProduct)
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

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "<h1>Hello world!")
}
