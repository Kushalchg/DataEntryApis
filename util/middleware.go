package util

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"www.github.com/kushalchg/DataEntryApis/global"
)

func GeneralAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header["Authorization"][0]
		// check if there is no authorization
		global.Logger.Printf("the authorization header value is %v\n", authorization)
		if authorization == "" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"error": "you forgot to add authorization header",
			})
			c.Abort()
			return

		}

		// check if the is token prefix "JWT" or not?
		tokenPrefix := strings.Split(authorization, " ")[0]
		if tokenPrefix != "JWT" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"error": "Must add JWT infront of token",
			})
			c.Abort()
			return
		}

		//check the token is valid or not
		token := strings.Split(authorization, " ")[1]
		global.Logger.Printf("the token value is %v\n", token)

		claims, err := ParseToken(token)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"error":  "Please provide valid token",
				"detail": err,
			})
			c.Abort()
			return
		}
		// check whethet the token is access or not?

		global.Logger.Printf("the category value is %v %v %v\n", claims.Cat, claims.Email, claims.Id)
		if claims.Cat != "access" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"error": "Please provide access token",
			})
			c.Abort()
			return
		}

		c.Next()
	}

}
