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

func (r *Rest) CreateFinancialTrx(c *gin.Context) {
	var (
		trx entity.FinancialTrx
	)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	var v entity.FinancialTrxReq
	if err := json.Unmarshal(body, &v); err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	trx, err = r.Uc.CreateFinancialTrx(c, v)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusCreated, trx)
}

func (r *Rest) UpdateFinancialTrx(c *gin.Context) {
	var (
		trx entity.FinancialTrx
	)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	var v entity.FinancialTrxReq
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

	trx, err = r.Uc.UpdateFinancialTrx(c, v)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusOK, trx)
}

func (r *Rest) GetFinancialTrxByID(c *gin.Context) {
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

	trx, err := r.Uc.GetFinancialTrxByID(c, vid)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusOK, trx)
}

func (r *Rest) DeleteFinancialTrxByID(c *gin.Context) {
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

	err = r.Uc.DeleteFinancialTrx(c, vid)
	if err != nil {
		r.HttpRespError(c, http.StatusBadRequest, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusOK, nil)
}
