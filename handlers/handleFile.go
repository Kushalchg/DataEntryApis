package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"www.github.com/kushalchg/DataEntryApis/global"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "please provide valid file",
			"detail": err.Error(),
		})
		return
	}
	// check file size must not exceed 10 MB
	if file.Size > 10<<20 {

		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "please provide valid file",
			"detail": "file size must not exceed 10MB",
		})
		return
	}
	global.Logger.Println(file.Filename)

	// Upload the file to specific dst.
	// dst := filepath.Join("uploadedFiles", file.Filename)
	c.SaveUploadedFile(file, fmt.Sprintf("uploadedFiles/%v", file.Filename))

	c.IndentedJSON(http.StatusOK, gin.H{

		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
	})
}

func GetFile(c *gin.Context) {
	filePath := "./uploadedFiles/6969.jpeg" // Specify the path to the file you want to serve
	c.FileAttachment(filePath, "6969.jpeg")
}
