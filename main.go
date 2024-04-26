package main

import (
	//"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/cookie"
	//"github.com/gin-gonic/gin"
	"github.com/robert/notification/model"
	"github.com/robert/notification/router"
)

func main() {
	model.Opendatabaseconnection()
	router.Startserver()
}
