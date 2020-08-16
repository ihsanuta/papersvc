package config

import "database/sql"

// type Config struct {
// }

// DBInit create connection to database
func DBInit() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/finance_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}
	return db
}
