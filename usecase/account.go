package usecase

import (
	"papersvc/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (u *Usecase) CreateAccount(c *gin.Context, v entity.ReqAccount) (entity.Account, error) {
	var account entity.Account
	payload := entity.Account{
		Username:  v.Username,
		Password:  v.Password,
		CreatedAt: time.Now(),
	}

	account, err := u.Domain.CreateSQLAccount(c, payload)
	if err != nil {
		return account, nil
	}

	return account, nil
}

func (u *Usecase) LoginAccount(c *gin.Context, v entity.ReqAccount) (string, error) {
	_, err := u.Domain.LoginSQLAccount(c, v)
	if err != nil {
		return "", err
	}

	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}
