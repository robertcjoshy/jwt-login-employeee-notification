package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID   uint   `json:"userid"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

var key = []byte("secret_key")

func CreateJWT(id uint, user string, c *gin.Context) string {

	expirationtime := time.Now().Add(time.Hour * 24)

	claims := &Claims{
		UserID:   id,
		UserName: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationtime.Unix(),
			Issuer:    "notification",
		},
	}
	fmt.Println(id, user)
	fmt.Println("----------------------")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // creating jwt token
	ss, err := token.SignedString(key)                         // signed using secretkey

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot sign token"})
	}
	return ss
}

func GetJWT(c *gin.Context) bool {

	//tokenSTRING := c.GetHeader("Authorization")
	tokenSTRING, _ := c.Cookie("jwt_token")
	fmt.Println(tokenSTRING)
	if tokenSTRING == "" {
		//c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		fmt.Println("NO VALUE IN HEADER")
		return false
	}
	fmt.Println("value in token = ", tokenSTRING)

	//token, err := jwt.ParseWithClaims(strings.SplitN(tokenSTRING, " ", 2)[1], &Claims{}, func(t *jwt.Token) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenSTRING, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	fmt.Println("tokenafterparsing", token)
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		fmt.Printf("header = %v , %v", claims.UserID, claims.UserName)
		return true
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
	}
	return false
}

func Isloggedin() gin.HandlerFunc {
	return func(c *gin.Context) {
		//tokenSTRING := c.Request.Header.Get("Authorization")
		tokenSTRING, _ := c.Cookie("jwt_token")
		if len(tokenSTRING) < 1 {
			fmt.Println("len = 0")
			c.Next()
			return
		}
		if tokenSTRING == "" {
			//c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Next()
			return
		}
		fmt.Println("no token")
		//token, err := jwt.ParseWithClaims(strings.SplitN(tokenSTRING, " ", 2)[1], &Claims{}, func(t *jwt.Token) (interface{}, error) {
		token, err := jwt.ParseWithClaims(tokenSTRING, &Claims{}, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		})

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			c.Set("username", claims.UserName)
			//c.Redirect(http.StatusOK, "/secret")
			c.Next()
		} else {
			//c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			fmt.Println("parsing error = ", err)
			c.Next()
		}
	}
}
