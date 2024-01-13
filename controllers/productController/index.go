package productcontroller

import (
	"net/http"

	"github.com/paafff/golang-CRUD/database"
	"github.com/paafff/golang-CRUD/models"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {

	var products []models.Product

	database.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return

	}

	database.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}
