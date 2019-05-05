package controllers

import (
    "github.com/gin-gonic/gin"
    "golang-backend/config"
    "golang-backend/entities"
	"net/http"
    "strconv"
)

var mSisa = 0

//JSONPagination : Pagination
func JSONPagination(c *gin.Context) {

    var mCount = 0
 
    Db, err := config.DbConnect()
	if err != nil {
		panic("Not Connect database")
    }
    
	var tbldft entities.DefaultPage
	
	row := Db.QueryRow("SELECT * FROM tabeldefaultpage;")
	err2 := row.Scan(&tbldft.ID, &tbldft.DefaultPage)	
	if err2 != nil {
		panic(err2)
    }
    
    mDefaultPage := tbldft.DefaultPage

    res := []entities.TabelCount{}
    var tblcnt entities.TabelCount
  
    row2 := Db.QueryRow("SELECT * FROM tabelcount;")
	err3 := row2.Scan(&tblcnt.ID,&tblcnt.CountProduct,&tblcnt.CountProductTemp)
	if err3 != nil {
		panic(err3)
	}
    res = append(res, tblcnt)
    Db.Close()

    param1 := c.Query("searchname")
    if param1 != "" {
        mCount = res[0].CountProductTemp 
    } else {
        mCount = res[0].CountProduct    
    }

    startend :=[]entities.StartEnd{}
    var tbl entities.StartEnd

    var moption = c.Query("moption")
    mStart, _ := strconv.Atoi(c.Query("mstart"))
    mEnd, _ := strconv.Atoi(c.Query("mend"))
    
    if moption == "0" {
        mStart = 0; mEnd = mDefaultPage
    }
    if moption == "1" {
        if mSisa != 0 {
            mStart = mStart-mDefaultPage ; mEnd = (mEnd-mDefaultPage)+(mDefaultPage-mSisa)
            mSisa = 0
        } else {
            mStart = mStart - mDefaultPage; mEnd = mEnd - mDefaultPage
        }
    }
    if moption == "2" {
        mStart = mStart + mDefaultPage; mEnd = mEnd + mDefaultPage
        if mStart >= mCount {
            mStart = mStart - mDefaultPage; mEnd = mEnd - mDefaultPage
        }
    }
    if moption == "3" {
        mSisa = mCount % mDefaultPage
        if mSisa != 0 {
            mStart = mCount - mSisa; mEnd = mCount
        } else {
            mStart = mCount - mDefaultPage; mEnd = mCount
        }
    }
    if mStart < 0 && mEnd <=0 {
       mStart = 0; mEnd = mDefaultPage
    }

    tbl.StartID = mStart
    tbl.EndID = mEnd
    startend = append(startend, tbl)
    c.JSON(http.StatusOK, startend)
}