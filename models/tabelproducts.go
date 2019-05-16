package models

import (
	// _ : import sql driver
	_ "github.com/go-sql-driver/mysql"
	"golang-pagination/entities"
	"database/sql"
)

// ListTblProduct : List tabelproduct per page
func ListTblProduct(db *sql.DB, tabelprd string) []entities.TabelProduct {

	res := []entities.TabelProduct{}
	var tbl entities.TabelProduct
	
	rows, err := db.Query(tabelprd)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	db.Close()
	for rows.Next() {
		err := rows.Scan(&tbl.ID, &tbl.ImagePath, &tbl.Title, &tbl.Description, &tbl.Price)
		if err != nil {
			panic(err)
		}
		res = append(res, tbl)
	}
	return res
}
