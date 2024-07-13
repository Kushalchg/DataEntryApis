package main

import (
	"www.github.com/kushalchg/DataEntryApis/initializers"
	"www.github.com/kushalchg/DataEntryApis/models"
)

func init() {
	initializers.LoadEnv()
	initializers.Connectdb()
}

func main() {
	//migrations
	initializers.DB.AutoMigrate(&models.User{})
}
