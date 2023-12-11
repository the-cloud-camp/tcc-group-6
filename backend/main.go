package main

import (
	AuthController "backend/controller/auth"
	ProductController "backend/controller/product"
	UserController "backend/controller/user"
	middleware "backend/middleware"
	"backend/orm"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	orm.InitDB()

	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
	r.GET("/products/:id", ProductController.GetAllProductsById)
	authorized := r.Group("", middleware.JWTAuthen())
	r.GET("/products/public", ProductController.GetAllProducts)
	authorized.GET("/products", ProductController.GetAllProductsByUser)
	authorized.POST("/products", ProductController.CreateProduct)
	authorized.DELETE("/products/:id", ProductController.DeleteProductsByID)

	authorizedusers := r.Group("/users", middleware.JWTAuthen())
	authorizedusers.GET("/readall", UserController.ReadAll)
	authorizedusers.GET("/profile", UserController.Profile)
	r.Run("localhost:8081")
}
