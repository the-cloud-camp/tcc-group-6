package matching

import (
	"backend/orm"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func SendOffer(c *gin.Context) {
	productIdSellString := c.Param("product_id_sell")
	productIdSell, SellErr := strconv.ParseUint(productIdSellString, 10, 64)

	productIdBuyString := c.Param("product_id_buy")
	productIdBuy, BuyErr := strconv.ParseUint(productIdBuyString, 10, 64)

	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("product_id_sell: %d, product_id_buy: %d\n", productIdSell, productIdBuy) // Debug Print Section
	fmt.Println(strings.Repeat("-", 100))

	if (SellErr != nil) || (BuyErr != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product_id"})
		return
	} else if productIdSell == productIdBuy {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product_id duplicated"})
		return
	}

	//===================== Search for existing products =====================//
	product1 := orm.Product{}
	product2 := orm.Product{}

	productSellResult := orm.Db.First(&product1, uint(productIdSell))
	productBuyResult := orm.Db.First(&product2, uint(productIdBuy))
	if (productSellResult.Error != nil) || (productBuyResult.Error != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}

	// c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "get product by id successfully"})

	//===================== Add new matching =====================//
	matching := orm.Matching{
		ProductIdBuy:  uint(productIdBuy),
		ProductIdSell: uint(productIdSell),
	}

	result := orm.Db.Create(&matching)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"status": "ERROR", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Matching created successfully", "matching": matching})
}

func GetOffer(c *gin.Context) {
	productIdSellString := c.Param("product_id_sell")
	productIdSell, SellErr := strconv.ParseUint(productIdSellString, 10, 64)

	productIdBuyString := c.Param("product_id_buy")
	productIdBuy, BuyErr := strconv.ParseUint(productIdBuyString, 10, 64)

	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("product_id_sell: %d, product_id_buy: %d\n", productIdSell, productIdBuy) // Debug Print Section
	fmt.Println(strings.Repeat("-", 100))

	if (SellErr != nil) || (BuyErr != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product_id"})
		return
	}

	matching := orm.Matching{}

	result := orm.Db.Where("product_id_sell = ? AND product_id_buy = ?", productIdSell, productIdBuy).First(&matching)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Matching not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Matching get successfully", "matching": matching})
}

func GetAllReceivedOffer(c *gin.Context) {
	userIdFloat64 := c.MustGet("userId").(float64)
	userId := uint(userIdFloat64)

	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("user_id: %d", userId) // Debug Print Section
	fmt.Println(strings.Repeat("-", 100))

	matching := []orm.Matching{}

	matchingResult := orm.Db.
		Table("matchings").
		Joins("JOIN products ON matchings.product_id_sell = products.id").
		Where("products.user_id = ? AND matchings.matching_status = 'available'", userId).
		Find(&matching)

	if matchingResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "AllReceivedMatching not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "AllReceivedMatching get successfully", "matching": matching})
}

func GetAllSentOffer(c *gin.Context) {
	userIdFloat64 := c.MustGet("userId").(float64)
	userId := uint(userIdFloat64)

	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("user_id: %d", userId) // Debug Print Section
	fmt.Println(strings.Repeat("-", 100))

	matching := []orm.Matching{}

	matchingResult := orm.Db.
		Table("matchings").
		Joins("JOIN products ON matchings.product_id_buy = products.id").
		Where("products.user_id = ? AND matchings.matching_status = 'available'", userId).
		Find(&matching)

	if matchingResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "AllSentMatching not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "AllSentMatching get successfully", "matching": matching})
}

func GetAllMatched(c *gin.Context) {
	userIdFloat64 := c.MustGet("userId").(float64)
	userId := uint(userIdFloat64)

	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("user_id: %d", userId) // Debug Print Section
	fmt.Println(strings.Repeat("-", 100))

	type Result struct {
		orm.Matching
		BuyerUserID    uint   `json:"buyer_user_id"`
		BuyerUsername  string `json:"buyer_username"`
		SellerUserID   uint   `json:"seller_user_id"`
		SellerUsername string `json:"seller_username"`
	}

	var matching []Result

	matchingResult := orm.Db.
		Table("matchings").
		Select("matchings.*, buy_products.user_id AS buyer_user_id, buy_products.username AS buyer_username, sell_products.user_id AS seller_user_id, sell_products.username AS seller_username").
		Joins("JOIN products AS buy_products ON matchings.product_id_buy = buy_products.id").
		Joins("JOIN products AS sell_products ON matchings.product_id_sell = sell_products.id").
		Where("(buy_products.user_id = ? OR sell_products.user_id = ?) AND matchings.matching_status = 'matched'", userId, userId).
		Find(&matching)

	if matchingResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "AllMatched not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "AllMatched get successfully", "matching": matching})
}

func GetMatchedInfo(c *gin.Context) {
	// matched := orm.Matching{}
	// matchedIdStr := c.Param("id")
	// matchedId, err := strconv.ParseUint(matchedIdStr, 10, 64)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "matched not found"})
	// 	return
	// }
	// result := orm.Db.Where("id = ?", matchedId).First(&matched)
	// if result.Error != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Matched not found"})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Get matched by id successfully", "matched": matched})

	userIdFloat64 := c.MustGet("userId").(float64)
	userId := uint(userIdFloat64)

	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("user_id: %d", userId) // Debug Print Section
	fmt.Println(strings.Repeat("-", 100))

	type Result struct {
		orm.Matching
		BuyerUserID     uint   `json:"buyer_user_id"`
		BuyerUsername   string `json:"buyer_username"`
		SellerUserID    uint   `json:"seller_user_id"`
		SellerUsername  string `json:"seller_username"`
		BuyerTelephone  string `json:"buyer_telephone"`
		SellerTelephone string `json:"seller_telephone"`
	}

	var matching Result

	matchingResult := orm.Db.
		Table("matchings").
		Select("matchings.*, buy_products.user_id AS buyer_user_id, buy_products.username AS buyer_username, sell_products.user_id AS seller_user_id, sell_products.username AS seller_username, buy_users.telephone AS buyer_telephone, sell_users.telephone AS seller_telephone").
		Joins("JOIN products AS buy_products ON matchings.product_id_buy = buy_products.id").
		Joins("JOIN products AS sell_products ON matchings.product_id_sell = sell_products.id").
		Joins("JOIN users AS buy_users ON buy_products.user_id = buy_users.id").
		Joins("JOIN users AS sell_users ON sell_products.user_id = sell_users.id").
		Where("(buy_products.user_id = ? OR sell_products.user_id = ?) AND matchings.matching_status = 'matched'", userId, userId).
		First(&matching)

	if matchingResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MatchingResult not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "MatchingResult get successfully", "matching": matching})
}
