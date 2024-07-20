package main

import (
	"time"

	"github.com/gin-contrib/cors"
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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // You can specify the allowed origins here
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.UserRoutes(r)
	routes.DataRoutes(r)
	routes.FileRoutes(r)
	r.Run()

}
