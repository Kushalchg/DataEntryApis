package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.github.com/kushalchg/DataEntryApis/handlers"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/users", handlers.UserRegister)
}
func DataRoutes(r *gin.Engine) {
	r.GET("/data", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"success": "ok form data route",
		})

	})
}
