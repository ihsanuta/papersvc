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
		r.HttpRespError(c, err)
		return
	}

	var v entity.FinancialTrxReq
	if err := json.Unmarshal(body, &v); err != nil {
		r.HttpRespError(c, err)
		return
	}

	trx, err = r.Uc.CreateFinancialTrx(c, v)
	if err != nil {
		r.HttpRespError(c, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusCreated, trx, nil)
}

func (r *Rest) UpdateFinancialTrx(c *gin.Context) {
	var (
		trx entity.FinancialTrx
	)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		r.HttpRespError(c, err)
		return
	}

	var v entity.FinancialTrxReq
	if err := json.Unmarshal(body, &v); err != nil {
		r.HttpRespError(c, err)
		return
	}

	idParam := c.Param("id")
	_, err = regexp.MatchString(`^[0-9]+$`, idParam)
	if err != nil {
		r.HttpRespError(c, err)
		return
	}

	v.ID, err = strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		r.HttpRespError(c, err)
		return
	}

	trx, err = r.Uc.UpdateFinancialTrx(c, v)
	if err != nil {
		r.HttpRespError(c, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusOK, trx, nil)
}

func (r *Rest) GetFinancialTrxByID(c *gin.Context) {
	idParam := c.Param("id")
	_, err := regexp.MatchString(`^[0-9]+$`, idParam)
	if err != nil {
		r.HttpRespError(c, err)
		return
	}

	vid, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		r.HttpRespError(c, err)
		return
	}

	trx, err := r.Uc.GetFinancialTrxByID(c, vid)
	if err != nil {
		r.HttpRespError(c, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusOK, trx, nil)
}

func (r *Rest) DeleteFinancialTrxByID(c *gin.Context) {
	idParam := c.Param("id")
	_, err := regexp.MatchString(`^[0-9]+$`, idParam)
	if err != nil {
		r.HttpRespError(c, err)
		return
	}

	vid, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		r.HttpRespError(c, err)
		return
	}

	err = r.Uc.DeleteFinancialTrx(c, vid)
	if err != nil {
		r.HttpRespError(c, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusOK, nil, nil)
}

func (r *Rest) GetFinancialTrx(c *gin.Context) {
	var p entity.FinancialTrxParam

	err := c.ShouldBindQuery(&p)
	if err != nil {
		r.HttpRespError(c, err)
	}

	result, pagination, err := r.Uc.GetFinancialTrx(c, p)
	if err != nil {
		r.HttpRespError(c, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusOK, result, pagination)
}
