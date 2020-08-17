package rest

import (
	"fmt"
	"net/http"
	"papersvc/usecase"
	"time"

	x "papersvc/errors"

	"papersvc/entity"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	Uc usecase.Usecase
}

func (r *Rest) HttpRespError(c *gin.Context, err error) {
	var meta entity.Meta
	statusCode, displayError := x.CompileError(err, "ID", false)
	statusStr := http.StatusText(statusCode)

	meta = entity.Meta{
		Path:       c.Request.RequestURI,
		StatusCode: statusCode,
		Status:     statusStr,
		Message:    fmt.Sprintf("%s %s [%d] %s", c.Request.Method, c.Request.RequestURI, statusCode, statusStr),
		Error:      &displayError,
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	c.JSON(statusCode, gin.H{
		"metadata": meta,
	})
}

func (r *Rest) HttpRespSuccess(c *gin.Context, status int, resp interface{}, pagination interface{}) {
	meta := entity.Meta{
		Path:       c.Request.RequestURI,
		StatusCode: status,
		Status:     http.StatusText(status),
		Message:    fmt.Sprintf("%s %s [%d] %s", c.Request.Method, c.Request.RequestURI, status, http.StatusText(status)),
		Error:      nil,
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	switch data := resp.(type) {
	case entity.Account:
		c.JSON(status, gin.H{
			"metadata": meta,
			"account":  &data,
		})
	case entity.FinancialAcc:
		c.JSON(status, gin.H{
			"metadata":     meta,
			"financialacc": &data,
		})
	case entity.FinancialTrx:
		c.JSON(status, gin.H{
			"metadata":     meta,
			"financialtrx": &data,
		})
	case []entity.FinancialAcc:
		c.JSON(status, gin.H{
			"metadata":     meta,
			"financialacc": &data,
			"pagination":   &pagination,
		})
	case []entity.FinancialTrx:
		c.JSON(status, gin.H{
			"metadata":     meta,
			"financialtrx": &data,
			"pagination":   &pagination,
		})
	case string:
		c.JSON(status, gin.H{
			"metadata": meta,
			"token":    &data,
		})
	case nil:
		c.JSON(status, gin.H{
			"metadata": meta,
		})
	default:
		r.HttpRespError(c, x.NewWithCode(x.CodeHTTPInternalServerError, fmt.Sprintf("cannot cast type of %+v", data)))
		return
	}
}
