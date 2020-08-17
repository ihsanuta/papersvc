package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		r.HttpRespError(c, err)
		return
	}

	var v entity.FinancialAccReq
	if err := json.Unmarshal(body, &v); err != nil {
		r.HttpRespError(c, err)
		return
	}

	account, err = r.Uc.CreateFinancialAcc(c, v)
	if err != nil {
		r.HttpRespError(c, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusCreated, account, nil)
}

func (r *Rest) UpdateFinancialAcc(c *gin.Context) {
	var (
		account entity.FinancialAcc
	)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		r.HttpRespError(c, err)
		return
	}

	var v entity.FinancialAccReq
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

	account, err = r.Uc.UpdateFinancialAcc(c, v)
	if err != nil {
		r.HttpRespError(c, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusOK, account, nil)
}

func (r *Rest) GetFinancialAccByID(c *gin.Context) {
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

	account, err := r.Uc.GetFinancialAccByID(c, vid)
	if err != nil {
		r.HttpRespError(c, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusOK, account, nil)
}

func (r *Rest) DeleteFinancialAccByID(c *gin.Context) {
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

	err = r.Uc.DeleteFinancialAcc(c, vid)
	if err != nil {
		r.HttpRespError(c, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusOK, nil, nil)
}

func (r *Rest) GetFinancialAcc(c *gin.Context) {
	var p entity.FinancialAccParam

	fmt.Printf("PARAM : %v", c.Request.URL.Query())

	err := c.ShouldBindQuery(&p)
	if err != nil {
		log.Fatal(err)
		r.HttpRespError(c, err)
	}

	fmt.Printf("PARAM RES : %v", p)

	result, pagination, err := r.Uc.GetFinancialAcc(c, p)
	if err != nil {
		log.Fatal(err)
		r.HttpRespError(c, err)
		return
	}

	r.HttpRespSuccess(c, http.StatusOK, result, pagination)
}
