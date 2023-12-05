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
	json := ProductBody{}
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

func GetAllProducts(c *gin.Context) {
	// products := []orm.Product{}
	// userIdFloat64 := c.MustGet("userId").(float64)
	// userId := uint(userIdFloat64)
	// if (userId == "") {
	// 	orm.Db.Find(&products)
	// } else {
	// 	orm.Db.Where("user_id!=?", userId).Find(&products)
	// }
	// c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "User read successfully", "products": products})
}

// func GetUserID(c *gin.Context) {
// 	userIdFloat64 := c.MustGet("userId").(float64)
// 	return userIdFloat64, err
// }
