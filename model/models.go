package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"gorm.io/gorm"
)

type User struct {
	Username string `gorm:"primarykey" json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (user *User) Save() error {

	err := Database.Model(&user).Create(&user).Error
	return err
}

func Getpassword(user string, c *gin.Context) (string, error) {
	var input User
	err := Database.Where("username = ?", user).First(&input).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid useranme"})
		c.Abort()
	}
	return input.Password, nil
}
