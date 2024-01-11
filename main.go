package main

import (
	"github.com/gin-gonic/gin"

	"github.com/paafff/golang-CRUD/models"
	"github.com/paafff/golang-CRUD/database"
	
)

func main()  {

	r := gin.Default()
	models.ConnectDatabase()

	
}