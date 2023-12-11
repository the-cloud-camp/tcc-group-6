package product

import (
	"backend/orm"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	product := orm.Product{
		Name:        json.Name,
		Description: json.Description,
		Price:       json.Price,
		Amount:      json.Amount,
		UserId:      userId,
		Username:    username,
	}

	result := orm.Db.Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"status": "ERROR", "message": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Product created successfully", "products": product})
}

func GetAllProducts(c *gin.Context) {
	products := []orm.Product{}
	orm.Db.Find(&products)
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Product get successfully", "products": products})
}

func GetAllProductsByUser(c *gin.Context) {
	products := []orm.Product{}
	userIdFloat64 := c.MustGet("userId").(float64)
	userId := uint(userIdFloat64)
	orm.Db.Where("user_id != ?", userId).Find(&products)
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Product get successfully", "products": products})
}

func GetAllProductsById(c *gin.Context) {
	product := orm.Product{}
	productIdStr := c.Param("id")
	productId, err := strconv.ParseUint(productIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}
	result := orm.Db.Where("id = ?", productId).First(&product)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "get product by id successfully", "product": product})
}

func DeleteProductsByID(c *gin.Context) {
	products := []orm.Product{}
	userIdFloat64 := c.MustGet("userId").(float64)
	userId := uint(userIdFloat64)
	productIdStr := c.Param("id")
	productId, err := strconv.ParseUint(productIdStr, 10, 64)

	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("user_id: %d, product_id: %d\n", userId, productId) // Debug Print Section
	fmt.Println(strings.Repeat("-", 100))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product_id"})
		return
	}

	result := orm.Db.Where("user_id = ? AND id = ?", userId, productId).First(&products)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	orm.Db.Where("user_id = ? AND id = ?", userId, productId).Delete(&products)
	// orm.Db.Where("user_id = ? AND id = ?", userId, productId).Unscoped().Delete(&products)
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Product deleted successfully", "products": products})
}
