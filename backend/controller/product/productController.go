package product

import (
	"backend/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ProductBody struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Amount      float64 `json:"amount"`
}

func CreateProduct(c *gin.Context) {
	var json ProductBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get UserId from token
	userIdFloat64 := c.MustGet("userId").(float64)
	username := c.MustGet("username").(string)
	userId := uint(userIdFloat64)
	product := orm.Product{Name: json.Name, Description: json.Description, Price: json.Price, Amount: json.Amount, UserId: userId, Username: username}

	orm.Db.Create(&product)
	if product.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Product created successfully", "product": product})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ERROR", "message": "Product created failed"})
	}
}
