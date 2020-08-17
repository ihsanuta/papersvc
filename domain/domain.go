package domain

import (
	"database/sql"
	"fmt"
	"papersvc/config"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

type Domain struct {
}

var DB *sql.DB

func init() {
	DB = config.DBInit()
}

func qBuilder(s interface{}, e string) (string, string, []interface{}, string) {
	where := "WHERE 1=1 "
	args := []interface{}{}
	sort := ""
	limit := ""
	page := ""

	rv := reflect.ValueOf(s)
	for i := 0; i < rv.NumField(); i++ {
		vField := rv.Field(i)
		vFieldItf := vField.Interface()
		var arg interface{}
		var lc int64 = 10

		switch f := vFieldItf.(type) {
		case int64:
			if f < 1 {
				continue
			}
			arg = f
		case string:
			if len(f) < 1 {
				continue
			}
			arg = f
		}

		col := rv.Type().Field(i).Tag.Get("param")
		field := rv.Type().Field(i).Tag.Get("db")

		if field != "sort" && field != "limit" && field != "page" {
			args = append(args, arg)
		}

		if field == "limit" {
			lc = arg.(int64)
		}

		switch col {
		case "sort":
			sort += fmt.Sprintf("ORDER BY %v ASC", arg)
		case "limit":
			limit += fmt.Sprintf("LIMIT %v", arg)
		case "page":
			paging := (arg.(int64) - 1) * lc
			page += fmt.Sprintf("OFFSET %v", paging)
		default:
			where += " AND " + field + " = ?"
		}
	}
	where += " AND " + e
	query := fmt.Sprintf("%s %s %s %s", where, limit, page, sort)
	return query, where, args, sort
}
