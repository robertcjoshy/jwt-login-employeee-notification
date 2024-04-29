package router

import (
	"github.com/gin-gonic/gin"
	"github.com/robert/notification/handler"
	"github.com/robert/notification/middlewares"
)

func Startserver() {

	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*.html")
	r.Use(middlewares.ErrorHandler) // centralized error handling

	// endpoints -------------
	r.GET("/", middlewares.Isloggedin(), handler.Homepage)
	r.GET("/login", handler.Loginget)
	r.POST("/login", handler.Loginpost)
	r.GET("/secret", handler.Secretpage)
	r.GET("/logout", handler.Logout)
	r.GET("/signup", handler.Signupget)
	r.POST("/signup", handler.Createuser)
	r.Run()

}
