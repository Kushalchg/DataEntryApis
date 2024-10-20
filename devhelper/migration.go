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
	// drop tables
	// initializers.DB.Migrator().DropTable(&models.User{}, &models.EntryData{})
	//migrations
	initializers.DB.AutoMigrate(&models.User{}, &models.EntryData{})
}
