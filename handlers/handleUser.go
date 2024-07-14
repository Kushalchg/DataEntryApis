package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.github.com/kushalchg/DataEntryApis/initializers"
	"www.github.com/kushalchg/DataEntryApis/models"
	"www.github.com/kushalchg/DataEntryApis/util"
)

func UserRegister(c *gin.Context) {

	var body struct {
		Email    string `json:"email" validate:"required,email"  `
		Password string `json:"password" validate:"required,min=8"`
	}

	if err := c.Bind(&body); err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"error":  "Error occured while register try again",
			"detail": err,
		})
		return
	}

	// // check the validataion
	// // email must be in email format
	// // password must contain min 8 letters
	// // conform password must match password
	// if err := global.Validate.Struct(&body); err != nil {
	// 	global.Logger.Printf("validation Failed %s \n", err)
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{
	// 		"error":  "required format is not met!",
	// 		"detail": err.Error(),
	// 	})
	// 	return
	// }

	// // create hash password
	// // hash password in []byte type
	// hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	// error occured on creating hash password
	// 	log.Fatal("error on creating hash password")
	// }

	// //It takes ConformPassword from user but doesn't upload ot the database
	// // ConformPassword is there to prevent user to enter unintended password.
	// user := types.User{Email: body.Email, Password: string(hashPassword)}
	user := models.User{Email: body.Email, Password: body.Password}
	util.SendMail()
	tx := initializers.DB.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "error occurred while registering user",
			"detail": err.Error(),
		})
		return
	}
	tx.Commit()

	// result := initializers.DB.Create(&user)
	// if result.Error != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{
	// 		"error":  "error occur while user register",
	// 		"detail": result.Error.Error(),
	// 	})
	// 	return
	// }

	c.IndentedJSON(http.StatusOK, gin.H{
		"value": "user created successfully",
		"data":  user,
	})

}
