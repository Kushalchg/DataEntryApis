package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"www.github.com/kushalchg/DataEntryApis/global"
	"www.github.com/kushalchg/DataEntryApis/initializers"
	"www.github.com/kushalchg/DataEntryApis/routes"
)

func init() {
	global.Validate = validator.New(validator.WithRequiredStructEnabled())
	initializers.LoadEnv()
	initializers.Connectdb()

}

func main() {
	r := gin.Default()
	r.Static("/converted/images", "./converted/images")

	//setting max file size
	r.MaxMultipartMemory = 5 << 20

	// cors setup
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// routes setup
	routes.UserRoutes(r)
	routes.DataRoutes(r)
	routes.FileRoutes(r)
	r.Run()

}
