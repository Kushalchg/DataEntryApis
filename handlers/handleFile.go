package handlers

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	dst := filepath.Join("uploadedFiles", file.Filename)
	// Upload the file to specific dst.
	// dst := "uploadedFiles/"
	c.SaveUploadedFile(file, dst)

	c.IndentedJSON(http.StatusOK, gin.H{

		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
	})
}
