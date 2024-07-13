package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"www.github.com/kushalchg/DataEntryApis/global"
	"www.github.com/kushalchg/DataEntryApis/initializers"
	"www.github.com/kushalchg/DataEntryApis/routes"
	"www.github.com/kushalchg/DataEntryApis/util"
)

func init() {
	global.Validate = validator.New(validator.WithRequiredStructEnabled())
	initializers.LoadEnv()
	initializers.Connectdb()

}

func main() {
	// initializers.DB.AutoMigrate(&models.User{})
	global.Logger = util.InitializeLogger()
	r := gin.Default()
	routes.UserRoutes(r)
	routes.DataRoutes(r)
	r.Run()

}
