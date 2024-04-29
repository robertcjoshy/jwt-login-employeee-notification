package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robert/notification/middlewares"
	"github.com/robert/notification/model"
)

func Homepage(c *gin.Context) {
	username, _ := c.Get("username")
	fmt.Println("user = ", username)
	c.Abort()
	fmt.Println("after abort")
	c.HTML(http.StatusOK, "home.html", gin.H{"username": username})
}

func Loginget(c *gin.Context) {
	result := middlewares.GetJWT(c)
	fmt.Println("gettoken = ", result)
	if !result {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	} else {
		c.Redirect(http.StatusFound, "/secret")
	}
}

func Loginpost(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	pass, err := model.Getpassword(username, c)

	if err == nil && pass == password {
		token := middlewares.CreateJWT(1, username, c)
		fmt.Println(token)
		//c.Header("Authorization", "Bearer "+token) // sending jwt token along with header
		c.SetCookie("jwt_token", token, 3600, "/", "", false, true)
		fmt.Println("sending header")
		//c.Redirect(http.StatusFound, "/secret")
		c.HTML(http.StatusOK, "secret.html", gin.H{})

	} else {
		errr := middlewares.NewAppError(http.StatusUnauthorized, "invalid credentials")
		//c.Error(errors.New("testing error"))
		c.Error(errr)
		return
	}
}

func Secretpage(c *gin.Context) {

	c.HTML(http.StatusFound, "secret.html", gin.H{})
}

func Logout(c *gin.Context) {
	//result := middlewares.GetJWT(c)
	c.SetCookie("jwt_token", "", -1, "/", "", false, true)
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func Signupget(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

func Createuser(c *gin.Context) {
	var input model.User

	// Bind form data to the user struct
	if err := c.Bind(&input); err != nil {
		errr := middlewares.NewAppError(http.StatusBadRequest, "cannot binf json")
		//c.JSON(http.StatusBadRequest, gin.H{"error": "cannot bind json"})
		c.Error(errr)
		return
	}
	// save to database
	err := input.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.HTML(http.StatusFound, "secret.html", gin.H{})
}
