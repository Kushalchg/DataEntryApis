package handlers

import (
	"fmt"
	"math/rand"
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
		fmt.Printf("validation Failed %s \n", err)
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

	// generating random intiger value with 6 digit and sent to the email for verification
	randomInt := 500000 + rand.Intn(20000)

	// //It takes ConformPassword from user but doesn't upload ot the database
	// // ConformPassword is there to prevent user to enter unintended password.

	user := models.User{Email: body.Email, Password: string(hashPassword), CoformCode: randomInt}

	// insert the data into table
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "error occur while user register",
			"detail": result.Error.Error(),
		})
		return
	}

	// send mail to user provided email address with conformation code
	err = util.SendMail([]string{body.Email}, []byte(fmt.Sprintf("Verify your email with %v", randomInt)))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":  "unable to send you email",
			"detail": err,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"value": "Thank you for registration",
		"data":  "Conformation mail is sent to your email",
	})

}

func UserLogin(c *gin.Context) {
	var user models.User
	var body struct {
		Email    string `json:"email" validate:"required,email"  `
		Password string `json:"password" validate:"required"`
	}

	// parse the response body data
	if err := c.Bind(&body); err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"error":  "Unable to bind body, try again",
			"detail": err,
		})
		return
	}
	// check whether the provided email is email format and password is not empty
	if err := global.Validate.Struct(&body); err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"error":  "Unable to validate, try again",
			"detail": err,
		})
		return
	}

	// retrieve the value from database with provided email
	result := initializers.DB.Where("email=?", body.Email).First(&user)

	if result.Error != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"error":  "Account with this credential doesnot exist",
			"detail": result.Error.Error(),
		})
		return
	}
	// check the provided password is correct or not by compare with database password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"data":   "Credential doesnot match",
			"result": "You probably provide wrong password ",
		})
		return
	}
	// generate access  jwt tokens

	fmt.Printf("the email and is value is%v %v \n ", body.Email, user.ID)
	accessToken, err := util.GenerateAccessToken(body.Email, "user", user.ID)

	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error":  "error while token generation",
			"detail": err.Error(),
		})
		return
	}

	// generate refresh jwt tokens
	refeshToken, err := util.GenerateRefreshToken(body.Email, "user", user.ID)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"error":  "error while token generation",
			"detail": err.Error(),
		})
		return
	}

	// provide the response

	c.IndentedJSON(http.StatusOK, gin.H{
		"detail": "Login Successful",
		"data": gin.H{
			"refresh": refeshToken,
			"access":  accessToken,
		},
	})
}

func VerifyUser(c *gin.Context) {

}

func UpdateProfile(c *gin.Context) {

}
func GetProfile(c *gin.Context) {

}
func GetRefresh(c *gin.Context) {

}
func UserLogout(c *gin.Context) {

}
