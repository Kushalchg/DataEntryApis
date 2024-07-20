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

	r.POST("/upload", handlers.UploadFile)
}
