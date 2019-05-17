package controllers

import (
	"golang-backend/entities"
	"golang-backend/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"encoding/json"
)

// CheckLogin : Cek user login
func CheckLogin(c *gin.Context) {

	var (
		tbl entities.TabelUser
		result  gin.H
	)

	Db, err := config.DbConnect()
	err = Db.Ping()
	if err != nil {
		result = gin.H {
			"code": 204,
			"status": "BAD",
        	"userlogin": "unknown",
        	"message": "Not Connect database",
		}
		Db.Close()
		c.JSON(http.StatusOK, result)
		return
	}

	data, e := ioutil.ReadAll(c.Request.Body)
	if e != nil {
		c.JSON(http.StatusBadRequest, e.Error())
	}

	var tabeluser entities.TabelUser
	e = json.Unmarshal(data, &tabeluser)
	username := tabeluser.UserName
	password := []byte(tabeluser.Password)
  
	mQuery:=`SELECT username, password, level,token from tabeluser WHERE username = '` + username + `';`
		
	rows, err := Db.Query(mQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	
	row := Db.QueryRow(mQuery)
	err = row.Scan(&tbl.UserName, &tbl.Password, &tbl.Level,&tbl.Token)
	if err != nil {
		result = gin.H {
			"code": 204,
			"status": "BAD",
        	"userlogin": "unknown",
        	"message": "Username not found",
		}
	} else {
		byteHash := []byte(tbl.Password)
		err := bcrypt.CompareHashAndPassword(byteHash, password)
		if err == nil {
			result = gin.H {
				"code": 200,
				"status": "OK",
				"userlogin": username,
				"passlogin": password,
				"token": tbl.Token,
				"level": tbl.Level,
				"message": "Login successful..",
			}
		} else {
			result = gin.H {
				"code": 204,
        		"status": "BAD",
        		"userlogin": "unknown",
        		"message": "Invalid password.. !",
			}
		}
	}
	Db.Close()
	c.JSON(http.StatusOK, result)
}
