package main

import (
	"fmt"
	"net/http"

	"github.com/ihsanuta/papersvc/rest"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	r := &rest.Rest{}

	// Account User
	router.POST("/v1/register", r.CreateAccount)
	router.POST("/v1/login", r.LoginAccount)

	// Financial Trx
	router.POST("/v1/trx", auth, r.CreateFinancialTrx)
	router.GET("/v1/trx", auth, r.GetFinancialTrx)
	router.GET("/v1/trx/:id", auth, r.GetFinancialTrxByID)
	router.PUT("/v1/trx/:id", auth, r.UpdateFinancialTrx)
	router.DELETE("/v1/trx/:id", auth, r.DeleteFinancialTrxByID)

	// Financial Account
	router.POST("/v1/account", auth, r.CreateFinancialAcc)
	router.GET("/v1/account", auth, r.GetFinancialAcc)
	router.GET("/v1/account/:id", auth, r.GetFinancialAccByID)
	router.PUT("/v1/account/:id", auth, r.UpdateFinancialAcc)
	router.DELETE("/v1/account/:id", auth, r.DeleteFinancialAccByID)

	router.Run(":3000")
}

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}
