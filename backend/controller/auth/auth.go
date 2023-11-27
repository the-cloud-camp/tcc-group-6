package auth

import (
	"backend/orm"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	key []byte
)

type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check User Exists
	var userExist orm.User
	orm.Db.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "ERROR", "message": "User already exists"})
		return
	}

	// Create User
	encryptPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)

	user := orm.User{Username: json.Username, Password: string(encryptPassword)}
	orm.Db.Create(&user)
	if user.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "User created successfully", "userId": user.ID})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ERROR", "message": "User created failed"})
	}
}

type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var json LoginBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check User Exists
	var userExist orm.User
	orm.Db.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "ERROR", "message": "User already exists"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(json.Password))
	if err == nil {
		key = []byte(os.Getenv("SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
			"userID": userExist.ID,
			"exp":    time.Now().Add(time.Minute * 1).Unix(),
		})

		tokenString, err := token.SignedString(key)
		fmt.Println(tokenString, err)

		c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Login successfully", "token": tokenString})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ERROR", "message": "Login failed"})
	}
}
