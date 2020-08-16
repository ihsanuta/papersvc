package rest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"papersvc/entity"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateFinancialAcc(c *gin.Context) {
	var (
		account entity.FinancialAcc
	)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	var v entity.FinancialAccReq
	if err := json.Unmarshal(body, &v); err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	account, err = r.Uc.CreateFinancialAcc(c, v)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusCreated, account)
}

func (r *Rest) UpdateFinancialAcc(c *gin.Context) {
	var (
		account entity.FinancialAcc
	)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	var v entity.FinancialAccReq
	if err := json.Unmarshal(body, &v); err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	idParam := c.Param("id")
	_, err = regexp.MatchString(`^[0-9]+$`, idParam)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	v.ID, err = strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	account, err = r.Uc.UpdateFinancialAcc(c, v)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusOK, account)
}

func (r *Rest) GetFinancialAccByID(c *gin.Context) {
	idParam := c.Param("id")
	_, err := regexp.MatchString(`^[0-9]+$`, idParam)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	vid, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	account, err := r.Uc.GetFinancialAccByID(c, vid)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusOK, account)
}

func (r *Rest) DeleteFinancialAccByID(c *gin.Context) {
	idParam := c.Param("id")
	_, err := regexp.MatchString(`^[0-9]+$`, idParam)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	vid, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	err = r.Uc.DeleteFinancialAcc(c, vid)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusOK, nil)
}
