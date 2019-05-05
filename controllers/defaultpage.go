package controllers

import (
    "github.com/gin-gonic/gin"
	"golang-backend/config"
	"golang-backend/entities"
	"net/http"
)
//JSONDefaultPage : Default Page
func JSONDefaultPage(c *gin.Context) {

    Db, err := config.DbConnect()
	if err != nil {
		panic("Not Connect database")
	}
	res := []entities.DefaultPage{}
	var tbldft entities.DefaultPage
	
	row := Db.QueryRow("SELECT * FROM tabeldefaultpage;")
	err2 := row.Scan(&tbldft.ID, &tbldft.DefaultPage)	
	if err2 != nil {
		panic(err2)
	}
	res = append(res, tbldft)
	Db.Close()
	c.JSON(http.StatusOK, res)
}