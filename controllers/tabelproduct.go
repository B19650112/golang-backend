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

var xParam1 string

// JSONListTblProduct : List seluruh tabelproduct
func JSONListTblProduct(c *gin.Context) {
	var sqlProduct, tabelCopy string

	Db, err := config.DbConnect()
	if err != nil {
		panic("Not Connect database")
	}

	tabelproduct := []entities.TabelProduct{}
	param1 := c.Query("searchname")
	param3 := c.Query("mstart")
	param4 := c.Query("mend")

	if param1 != "" {
		if param1 != xParam1 {
			xParam1 = param1
			DelTempProduct()  // Delete producttemp before
			tabelCopy = `Call copy_tabelproduct('`+param1+`');` // Copy to producttemp for searchname product name
			rows, err := Db.Query(tabelCopy)
			if err != nil {
				panic(err)
			}
			defer rows.Close()
			CountProductTemp(param1)  // Update countproducttemp of tabelcount
		}
	} else {
		DelTempProduct()  // Delete producttemp
	}

	if param1 != "" {
		sqlProduct = `SELECT * FROM producttemp WHERE idtemp > '`+param3+`' AND idtemp <='`+param4+`' ORDER By idtemp;`
	} else {
		sqlProduct = `SELECT * FROM tabelproduct WHERE id > '`+param3+`' AND id <='`+param4+`' ORDER By id;`
	}
	
	dataList := models.ListTblProduct(Db, sqlProduct)
	tabelproduct = dataList
	c.JSON(http.StatusOK, tabelproduct)
	tabelproduct = nil
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
	dataList := models.ViewTblProduct(Db, sqlProduct)
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

	var numRows = 1
	var mQuery2 = `Call update_countproduct('` + strconv.Itoa(numRows) + `');`
	updDB2, err := Db.Query(mQuery2)
	if err != nil {
		panic(err)
	}
	defer updDB2.Close()

	Db.Close()
	c.JSON(http.StatusOK, "Add Record Success")
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
	c.JSON(http.StatusOK, "Update Record Success")
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

	var numRows = -1
	var mQuery2 = `Call update_countproduct('` + strconv.Itoa(numRows) + `');`
	updDB2, err := Db.Query(mQuery2)
	if err != nil {
		panic(err)
	}
	defer updDB2.Close()
	Db.Close()
	c.JSON(http.StatusOK, "Delete Record Success")
}
//JSONCheckTblProduct : input validation tabelproduct
func JSONCheckTblProduct(c *gin.Context) {

	imagepath := c.Query("imagepath")
	title := c.Query("title")
	description := c.Query("description")
	price := c.Query("price")
	mfound := c.Query("mfound")

	if imagepath == "" {
		c.JSON(http.StatusOK, "Path Image harus diisi")
		return
	}
	if title == "" {
		c.JSON(http.StatusOK, "Title harus diisi")
		return
	}
	if description == "" {
		c.JSON(http.StatusOK, "Description harus diisi")
		return
	}
	if price == "" || price == "0" {
		c.JSON(http.StatusOK, "Price harus > 0")
		return
	}
	if mfound == "1" {
		c.JSON(http.StatusOK, "")
		return
	}
	c.JSON(http.StatusOK, "")
}
// CountProductTemp : Update countproducttemp of tabelcount
func CountProductTemp(mSearch string) {
	Db, err := config.DbConnect()
	if err != nil {
		panic("Not Connect database")
	}
	var countProduct int
	if mSearch != "" {
		mQuery := `Call count_tabelproduct('` + mSearch + `');`
		err := Db.QueryRow(mQuery).Scan(&countProduct)
		if err != nil {
			panic(err)
		}
		var mQuery2 = `Call update_countproducttemp('` + strconv.Itoa(countProduct) + `');`
		updDB2, err := Db.Query(mQuery2)
		if err != nil {
			panic(err)
		}
		defer updDB2.Close()
	}
	Db.Close()
}
// DelTempProduct : Delete temporary file 
func DelTempProduct() {
	Db, err := config.DbConnect()
	if err != nil {
		panic("Not Connect database")
	}
	mQuery := `DROP TABLE IF EXISTS producttemp;`
	rows, err := Db.Query(mQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	Db.Close()
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
