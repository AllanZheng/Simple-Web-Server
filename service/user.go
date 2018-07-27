package service

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreteUser(c *gin.Context){
	c.HTML(http.StatusOK,"create.html",gin.H{
	})
}
func UpdateUser(c *gin.Context){
    c.HTML(http.StatusOK,"",gin.H{})
}

func SearchUser(c *gin.Context){
	c.HTML(http.StatusOK,"",gin.H{})
}


func DeleteUser(c *gin.Context){
	c.JSON(http.StatusOk,"",gin.H{})
}