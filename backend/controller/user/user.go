package user

import (
	"backend/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadAll(c *gin.Context) {
	var users []orm.User
	orm.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "User read successfully", "users": users})
}

func Profile(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var users []orm.User
	orm.Db.Find(&users, userId)
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "User read successfully", "users": users})
}