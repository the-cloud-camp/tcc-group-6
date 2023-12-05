package main

import (
	AuthController "backend/controller/auth"
	UserController "backend/controller/user"
	ProductController "backend/controller/product"
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
	authorized := r.Group("", middleware.JWTAuthen())
	authorized.POST("/products", ProductController.CreateProduct)
	// r.GET("/products", ProductController.GetProduct)

	authorizedusers := r.Group("/users", middleware.JWTAuthen())
	authorizedusers.GET("/readall", UserController.ReadAll)
	authorizedusers.GET("/profile", UserController.Profile)
	r.Run("localhost:8081")
}
