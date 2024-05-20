package main

import (
	"github.com/gin-gonic/gin"

	productcontroller "github.com/paafff/golang-CRUD/controllers/productController"
	"github.com/paafff/golang-CRUD/database"
)

func main() {

	r := gin.Default()
	database.ConnectDatabase()

	r.GET("/products", productcontroller.GetProducts)
	r.POST("/products", productcontroller.CreateProduct)

	r.Run(":5000")

}
