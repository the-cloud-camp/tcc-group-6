package main

import (
	AuthController "backend/controller/auth"
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
	authorized := r.Group("/users", middleware.JWTAuthen())
	authorized.GET("/readall", UserController.ReadAll)
	authorized.GET("/profile", UserController.Profile)
	r.Run("localhost:8081")
}
