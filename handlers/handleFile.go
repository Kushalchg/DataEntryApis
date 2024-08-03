package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"www.github.com/kushalchg/DataEntryApis/global"
	"www.github.com/kushalchg/DataEntryApis/util"
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
	actuaFileName := time.Now().UnixMicro()
	// actuaFileName := uuid.New().String()

	c.SaveUploadedFile(file, fmt.Sprintf("uploadedFiles/%v", actuaFileName))
	htmlFile, textFile, imageFile := util.AsciiConverter(fmt.Sprintf("uploadedFiles/%v", actuaFileName))
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		"html":    htmlFile,
		"text":    textFile,
		"image":   imageFile,
	})
}

// func GetFile(c *gin.Context) {
// 	fmt.Println("GetFile handler called")

// 	// Log the request method and content type
// 	fmt.Printf("Request Method: %s\n", c.Request.Method)
// 	fmt.Printf("Content-Type: %s\n", c.GetHeader("Content-Type"))

// 	// Log query parameters
// 	fmt.Printf("Query parameters: %v\n", c.Request.URL.Query())

// 	// Log request body
// 	bodyData, _ := io.ReadAll(c.Request.Body)
// 	fmt.Printf("Request body: %s\n", string(bodyData))
// 	// Restore the body so it can be read again in binding
// 	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyData))

// 	var body struct {
// 		FileName string `json:"fileName" form:"fileName" uri:"fileName"`
// 	}

// 	// Try different binding methods
// 	if err := c.ShouldBindJSON(&body); err == nil {
// 		fmt.Println("Bound using JSON")
// 	} else if err := c.ShouldBindQuery(&body); err == nil {
// 		fmt.Println("Bound using Query")
// 	} else if err := c.ShouldBindUri(&body); err == nil {
// 		fmt.Println("Bound using URI")
// 	} else if err := c.ShouldBind(&body); err == nil {
// 		fmt.Println("Bound using generic Bind")
// 	} else {
// 		fmt.Printf("All binding methods failed: %v\n", err)
// 	}

// 	fmt.Printf("FileName after binding: %s\n", body.FileName)

// 	if body.FileName == "" {
// 		fmt.Println("FileName is empty")
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "fileName is required",
// 		})
// 		return
// 	}

// 	filePath := filepath.Join("./output", body.FileName)
// 	fmt.Printf("File path: %s\n", filePath) // Debug print

// 	// Check if file exists
// 	if _, err := os.Stat(filePath); os.IsNotExist(err) {
// 		fmt.Println("File not found") // Debug print
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"error": "file not found",
// 		})
// 		return
// 	}

//		fmt.Println("Serving file") // Debug print
//		c.FileAttachment(filePath, body.FileName)
//	}
func GetHtmlFile(c *gin.Context) {

	var body struct {
		FileName string `json:"fileName" form:"fileName" uri:"fileName"`
	}

	bodyData, _ := io.ReadAll(c.Request.Body)

	// Restore the body so it can be read again in binding
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyData))

	// Try different binding methods
	if err := c.ShouldBindJSON(&body); err == nil {
		fmt.Println("Bound using JSON")
	}

	if body.FileName == "" {
		fmt.Println("FileName is empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "fileName is required",
		})
		return
	}

	filePath := filepath.Join("./converted/html/", body.FileName)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File not found") // Debug print
		c.JSON(http.StatusNotFound, gin.H{
			"error": "file not found",
		})
		return
	}

	c.FileAttachment(filePath, body.FileName)
}

func GetTextFile(c *gin.Context) {

	var body struct {
		FileName string `json:"fileName" form:"fileName" uri:"fileName"`
	}

	bodyData, _ := io.ReadAll(c.Request.Body)

	// Restore the body so it can be read again in binding
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyData))

	// Try different binding methods
	if err := c.ShouldBindJSON(&body); err == nil {
		fmt.Println("Bound using JSON")
	}

	if body.FileName == "" {
		fmt.Println("FileName is empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "fileName is required",
		})
		return
	}

	filePath := filepath.Join("./converted/text/", body.FileName)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File not found") // Debug print
		c.JSON(http.StatusNotFound, gin.H{
			"error": "file not found",
		})
		return
	}

	c.FileAttachment(filePath, body.FileName)
}

func GetImageFile(c *gin.Context) {

	var body struct {
		FileName string `json:"fileName" form:"fileName" uri:"fileName"`
	}

	bodyData, _ := io.ReadAll(c.Request.Body)

	// Restore the body so it can be read again in binding
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyData))

	// Try different binding methods
	if err := c.ShouldBindJSON(&body); err == nil {
		fmt.Println("Bound using JSON")
	}

	if body.FileName == "" {
		fmt.Println("FileName is empty")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "fileName is required",
		})
		return
	}

	filePath := filepath.Join("./converted/images/", body.FileName)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File not found") // Debug print
		c.JSON(http.StatusNotFound, gin.H{
			"error": "file not found",
		})
		return
	}

	c.FileAttachment(filePath, body.FileName)
}
