package main

import (
	"jwt/config"
	"jwt/controllers"
	"jwt/midleware"

	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectToDb()
}
func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.Signin)
	r.POST("/addproduct", midleware.ReqAuth, controllers.Addproduct)
	r.GET("/getallproducets", controllers.Getproducts)
	r.GET("/getuserproducts/:id",controllers.GetUserProductes)
	r.Run()
}
