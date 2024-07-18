package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.github.com/kushalchg/DataEntryApis/global"
	"www.github.com/kushalchg/DataEntryApis/initializers"
	"www.github.com/kushalchg/DataEntryApis/models"
)

func InsertEntryData(c *gin.Context) {
	var body struct {
		Tname     string  `validate:"required" json:"tname"`
		Tlength   float32 `validate:"required" json:"tlength"`
		Tdiameter float32 `josn:"tdiameter"`
		Tlogi     float32 `josn:"tlongi"`
		Tlatt     float32 `josn:"tlatt"`
	}
	// bind the request data into body
	if err := c.Bind(&body); err != nil {
		global.Logger.Printf("binding request body Failed %s \n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "required format is not met",
			"detail": err,
		})
		return
	}
	// validate the struct
	if err := global.Validate.Struct(&body); err != nil {
		global.Logger.Printf("validation Failed %s \n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "validation failed",
			"detail": err,
		})

		return

	}

	dataEntry := models.EntryData{Tname: body.Tname, Tlength: body.Tlength, Tdiameter: body.Tdiameter, Tlogi: body.Tlogi, Tlatt: body.Tlatt, UId: 1}
	result := initializers.DB.Create(&dataEntry)

	if result.Error != nil {
		global.Logger.Printf("data insert  Failed %s \n", result.Error.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "Problom occur while inserting data",
			"detail": result.Error.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"value": "data inserted successfully",
	})

}
