package models

import (
	// _ : import sql driver
	_ "github.com/go-sql-driver/mysql"
	"golang-backend/entities"
	"database/sql"
)

// ListTblResume : List tabelresume all
func ListTblResume(db *sql.DB, sql string) []entities.TabelResume {

	res := []entities.TabelResume{}
	var tbl entities.TabelResume
	
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	db.Close()

	for rows.Next() {
		err = rows.Scan(&tbl.ID, &tbl.Name, &tbl.Title, &tbl.Photo, &tbl.Description)
		if err != nil {
			panic(err)
		}
		res = append(res, tbl)
	}
	return res
}
