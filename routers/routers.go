package routers

import (
	"golang-pagination/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

//SetupRouter : setup router by group
func SetupRouter() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())

	grp00 := r.Group("api/grp00")
	{
		grp00.POST("/ceklogin", controllers.CheckLogin)
	}
	
	grp01 := r.Group("api/grp01")
	{
		grp01.GET("/listtblproduct", controllers.JSONListTblProduct)
		grp01.POST("/addtblproduct", controllers.JSONAddTblProduct)
		grp01.POST("/updatetblproduct/:id", controllers.JSONUpdateTblProduct)
		grp01.DELETE("/deletetblproduct/:id", controllers.JSONDeleteTblProduct)
		grp01.GET("/checktblproduct", controllers.JSONCheckTblProduct)
	}

	grp02 := r.Group("api/grp02")
	{ 
		 grp02.GET("/listtblresume", controllers.JSONListTblResume)
	}

	return r
}
