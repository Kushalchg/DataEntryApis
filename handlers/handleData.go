package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"www.github.com/kushalchg/DataEntryApis/global"
	"www.github.com/kushalchg/DataEntryApis/initializers"
	"www.github.com/kushalchg/DataEntryApis/models"
	"www.github.com/kushalchg/DataEntryApis/util"
)

func InsertEntryData(c *gin.Context) {
	var body struct {
		Tname     string  `validate:"required" json:"tname"`
		Tlength   float32 `validate:"required" json:"tlength"`
		Tdiameter float32 `josn:"tdiameter"`
		Tlongi    float32 `josn:"tlongi"`
		Tlatt     float32 `josn:"tlatt"`
	}

	// bind the request data into body
	if err := c.Bind(&body); err != nil {
		fmt.Printf("binding request body Failed %s \n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "required format is not met",
			"detail": err,
		})
		return
	}

	// validate the struct
	if err := global.Validate.Struct(&body); err != nil {
		fmt.Printf("validation Failed %s \n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "validation failed",
			"detail": err,
		})
		return
	}
	//get token from request header and parse and get the id of user
	authorization := c.Request.Header["Authorization"][0]
	claims, err := util.ParseToken(strings.Split(authorization, " ")[1])
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "Technical Error occur",
			"detail": err,
		})
		return
	}
	fmt.Printf("the authorixation header is %v\n", authorization)

	dataEntry := models.EntryData{Tname: body.Tname, Tlength: body.Tlength, Tdiameter: body.Tdiameter, Tlongi: body.Tlongi, Tlatt: body.Tlatt, UId: int(claims.Id)}
	result := initializers.DB.Create(&dataEntry)

	if result.Error != nil {
		fmt.Printf("data insert  Failed %s \n", result.Error.Error())
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

func DeleteData(c *gin.Context) {

}

func UpdateData(c *gin.Context) {

}

func GetAllData(c *gin.Context) {

}
func GetSingleData(c *gin.Context) {

}
