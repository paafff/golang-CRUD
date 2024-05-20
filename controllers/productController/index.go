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

	// Menghandle upload gambar
	file, err := c.FormFile("productImage")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengunggah gambar"})
		return
	}

	// Mendapatkan nama dan URL gambar
	product.ProductImageName = file.Filename
	product.ProductImageURL = "http://localhost:5000/assets/productImage/" + file.Filename

	// Menyimpan file gambar di direktori server
	dst := "./assets/productImage" + file.Filename
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menyimpan gambar"})
		return
	}

	database.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}
