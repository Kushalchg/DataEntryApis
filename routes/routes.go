package routes

import (
	"github.com/gin-gonic/gin"
	"www.github.com/kushalchg/DataEntryApis/handlers"
	"www.github.com/kushalchg/DataEntryApis/util"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/user/register", handlers.UserRegister)
	r.POST("/user/login", handlers.UserLogin)
}
func DataRoutes(r *gin.Engine) {
	r.POST("/insert-data", util.GeneralAuth(), handlers.InsertEntryData)
}
func FileRoutes(r *gin.Engine) {

	file := r.Group("/file")
	{
		file.POST("/upload", handlers.UploadFile)
		file.GET("/download/html", handlers.GetHtmlFile)
		file.GET("/download/text", handlers.GetTextFile)
		file.GET("/download/image", handlers.GetImageFile)
	}
}
