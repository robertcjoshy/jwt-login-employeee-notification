package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (user *User) Save() error {

	err := Database.Model(&user).Create(&user).Error
	return err
}

func Getpassword(user string, c *gin.Context) (User, error) {
	var input User
	err := Database.Where("id = ?", user).First(&input).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid useranme"})
		c.Abort()
	}
	return input, nil
}
