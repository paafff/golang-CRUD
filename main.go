package main

import (
	"github.com/gin-gonic/gin"


	"github.com/paafff/golang-CRUD/database"
	"github.com/paafff/golang-CRUD/controllers/productController"
	
)

func main()  {

	r := gin.Default()
	database.ConnectDatabase()

r.GET("/products", productcontroller.GetProducts)

	r.Run(":5000")
	
}