package controllers

import (
	"golang-backend/config"
	"golang-backend/entities"
	"golang-backend/models"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"io/ioutil"
)

// JSONListTblProduct : List all tabelproduct
func JSONListTblProduct(c *gin.Context) {
	var sqlProduct string

	Db, err := config.DbConnect()
	if err != nil {
		panic("Not Connect database")
	}

	tabelproduct := []entities.TabelProduct{}

	parmSearch := c.Query("searchname")
	if parmSearch != "" {
		idSearch := "%"+parmSearch+"%"
		sqlProduct = `SELECT * FROM tabelproduct where description like '`+ idSearch+`' order by id;`
	} else {
		sqlProduct = `SELECT * FROM tabelproduct order by id;`
	}
	
	dataList:= models.ListTblProduct(Db, sqlProduct)
	tabelproduct = dataList
	Db.Close()

	c.JSON(http.StatusOK, tabelproduct)
	
}
// JSONViewTblProduct : view or edit display tabelproduct by id
func JSONViewTblProduct(c *gin.Context) {
	Db, err := config.DbConnect()
	if err != nil {
		panic("Not Connect database")
	}
	tabelproduct := []entities.TabelProduct{}
	paramID := c.Param("id")
	sqlProduct := `SELECT * FROM tabelproduct WHERE id ='` + paramID + `';`
	dataList := models.ListTblProduct(Db, sqlProduct)
	tabelproduct = dataList
	c.JSON(http.StatusOK, tabelproduct)
	tabelproduct = nil
}
// JSONAddTblProduct : Add tabelproduct
func JSONAddTblProduct(c *gin.Context) {
	Db, err := config.DbConnect()
	if err != nil {
		panic("Not Connect database")
	}

	data, e := ioutil.ReadAll(c.Request.Body)
	if e != nil {
		c.JSON(http.StatusBadRequest, e.Error())
	}
	var tabelproduct entities.TabelProduct
	err2 := json.Unmarshal(data, &tabelproduct)
	if err2 != nil {
	 	fmt.Println(err2)
	}
	
	imagepath := tabelproduct.ImagePath
	title := tabelproduct.Title
	description := tabelproduct.Description
	price := strconv.Itoa(tabelproduct.Price)
	mQuery := `INSERT INTO tabelproduct SET imagepath ='` + imagepath + `', title='` + title + `', description='` + description + `',
				price='` + price + `';`
	updDB, err := Db.Query(mQuery)
	if err != nil {
		panic(err)
	}
	defer updDB.Close()

	Db.Close()
	c.JSON(http.StatusOK, "Add record successfully")
}
// JSONUpdateTblProduct : Update tabelproduct
func JSONUpdateTblProduct(c *gin.Context) {
	Db, err := config.DbConnect()
	if err != nil {
		panic("Not Connect database")
	}
	paramID := c.Param("id")

	data, e := ioutil.ReadAll(c.Request.Body)
	if e != nil {
		c.JSON(http.StatusBadRequest, e.Error())
	}

	var tabelproduct entities.TabelProduct
	err2 := json.Unmarshal(data, &tabelproduct)
	if err2 != nil {
	 	fmt.Println(err2)
	}
	
	imagepath := tabelproduct.ImagePath
	title := tabelproduct.Title
	description := tabelproduct.Description
	price := strconv.Itoa(tabelproduct.Price)
	mQuery := `UPDATE tabelproduct SET imagepath ='` + imagepath + `', title='` + title + `', description='` + description + `',
					  price='` + price + `' WHERE id ='`+paramID+`';`							  
	updDB, err := Db.Query(mQuery)
	if err != nil {
		panic(err)
	}
	defer updDB.Close()
	Db.Close()
	c.JSON(http.StatusOK, "Update record successfully")
}
// JSONDeleteTblProduct : Delete tabelproduct
func JSONDeleteTblProduct(c *gin.Context) {
	Db, err := config.DbConnect()
	if err != nil {
		panic("Not Connect database")
	}
	paramID := c.Param("id")
	query := `DELETE FROM tabelproduct WHERE id ='` + paramID + `';`
	updDB, err := Db.Query(query)
	if err != nil {
		panic(err)
	}
	defer updDB.Close()
	Db.Close()
	c.JSON(http.StatusOK, "Delete record successfully")
}
//JSONCheckTblProduct : input validation tabelproduct
func JSONCheckTblProduct(c *gin.Context) {

	imagepath := c.Query("imagepath")
	title := c.Query("title")
	description := c.Query("description")
	price := c.Query("price")

	if imagepath == "" {
		c.JSON(http.StatusOK, "Image Path can't blank !")
		return
	}
	if title == "" {
		c.JSON(http.StatusOK, "Title can't blank !")
		return
	}
	if description == "" {
		c.JSON(http.StatusOK, "Description can't blank !")
		return
	}
	if price == "" || price == "0" {
		c.JSON(http.StatusOK, "Price must greater than 0")
		return
	}
	c.JSON(http.StatusOK, "")
}
//JSONCountTblProduct : Search count tabelproduct
func JSONCountTblProduct(c *gin.Context) {
	Db, err := config.DbConnect()
	if err != nil {
		panic("Not Connect database")
	}
	var countProduct int
	tblcount := `SELECT countproduct FROM tabelcount;`
	err2 := Db.QueryRow(tblcount).Scan(&countProduct)
	if err2 != nil {
		panic(err2)
	}
	Db.Close()
	c.JSON(http.StatusOK, countProduct)
}
// JSONPrintTblProduct : Print tabelproduct
func JSONPrintTblProduct(c *gin.Context) {
	Db, err := config.DbConnect()
	if err != nil {
		panic("Not Connect database")
	}
	tabelproduct := []entities.TabelProduct{}
	sqlProduct := `SELECT * FROM tabelproduct;`
	dataList := models.ListTblProduct(Db, sqlProduct)
	tabelproduct = dataList
	Db.Close()
	c.JSON(http.StatusOK, tabelproduct)
}
