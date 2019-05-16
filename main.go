package main

import (
	"golang-pagination/routers"
	"database/sql"
)

// Db : type Database
var Db *sql.DB

func main() {

	r := routers.SetupRouter()
	//running
	r.Run(":8080")
}
