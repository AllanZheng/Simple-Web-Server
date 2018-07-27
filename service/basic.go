package service

import(
	"github.com/gin-gonic/gin"

	"net/http"
)

func LoginForm(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"OK!",
	})
}


func Logout(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"OK!",
	})
}

func PostLoginForm(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"OK!",
	})
}

func HomePage(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"OK!",
	})
}