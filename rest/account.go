package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"papersvc/entity"

	"github.com/gin-gonic/gin"
)

// create new data to the database
func (r *Rest) CreateAccount(c *gin.Context) {
	var (
		account entity.Account
	)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	var v entity.ReqAccount
	if err := json.Unmarshal(body, &v); err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	account, err = r.Uc.CreateAccount(c, v)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusCreated, account)
}

func (r *Rest) LoginAccount(c *gin.Context) {
	var (
		token string
	)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	var v entity.ReqAccount
	if err := json.Unmarshal(body, &v); err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	token, err = r.Uc.LoginAccount(c, v)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusOK, token)
}
