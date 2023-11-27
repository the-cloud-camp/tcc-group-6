package auth

import (
	"backend/orm"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
