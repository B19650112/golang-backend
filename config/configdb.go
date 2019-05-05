package config

import (
	"database/sql"
	"fmt"
)

// DbConnect : Connect to Database MySql
func DbConnect() (*sql.DB, error) {
	conn, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/mysqldata?parseTime=true")
	if err != nil {
		fmt.Println("gagal..!")
		return nil, err
	}
	return conn, nil
}
