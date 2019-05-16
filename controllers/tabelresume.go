package controllers

import (
	"golang-pagination/entities"
	"golang-pagination/config"
	"golang-pagination/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

// JSONListTblResume : List seluruh tabelResume
func JSONListTblResume(c *gin.Context) {

	Db, err := config.DbConnect()
	if err != nil {
		fmt.Println(err)
		panic("Not Connect database")
	}

	tabelresume := []entities.TabelResume{}	
	sqlResume := `SELECT * FROM tabelresume;`	
	dataList := models.ListTblResume(Db, sqlResume)
	tabelresume = dataList
	
	c.JSON(http.StatusOK, tabelresume)
	tabelresume = nil
}
