package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"www.github.com/kushalchg/DataEntryApis/global"
	"www.github.com/kushalchg/DataEntryApis/initializers"
	"www.github.com/kushalchg/DataEntryApis/models"
	"www.github.com/kushalchg/DataEntryApis/util"
)

func UserRegister(c *gin.Context) {

	var body struct {
		Email           string `json:"email" validate:"required,email"  `
		Password        string `json:"password" validate:"required,min=8"`
		ConformPassword string `json:"conform_password" validate:"required,eqfield=Password"`
	}

	if err := c.Bind(&body); err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"error":  "Error occured while register try again",
			"detail": err,
		})
		return
	}

	// check the validataion
	// email must be in email format
	// password must contain min 8 letters
	// conform password must match password
	if err := global.Validate.Struct(&body); err != nil {
		global.Logger.Printf("validation Failed %s \n", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "required format is not met!",
			"detail": err.Error(),
		})
		return
	}

	// create hash password
	// hash password in []byte type
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		// error occured on creating hash password
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "error occurred while registering user",
			"detail": err.Error(),
		})
		return

	}

	// //It takes ConformPassword from user but doesn't upload ot the database
	// // ConformPassword is there to prevent user to enter unintended password.
	user := models.User{Email: body.Email, Password: string(hashPassword)}

	// insert the data into table
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "error occur while user register",
			"detail": result.Error.Error(),
		})
		return
	}
	hello := 2000
	err = util.SendMail([]string{body.Email}, []byte(fmt.Sprintf("Verify your email with %v", hello)))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "unable to send you email",
			"detail": err,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"value": "user created successfully",
		"data":  "the mail is sent to your email",
	})

}
