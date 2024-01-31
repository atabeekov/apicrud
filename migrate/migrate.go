package main

import (
	"Api-CRUD/initializers"
	"Api-CRUD/models"
)

func init() {
	initializers.LoadVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
