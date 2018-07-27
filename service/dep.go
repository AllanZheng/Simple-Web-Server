package service

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateDep(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"FAILED!",
	})
}
func UpdateDep(c *gin.Context){
    c.JSON(http.StatusOK,gin.H{

	}})
}

func SearchDep(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{

	})
}


func DeleteDep(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{

	})
}

func SearchAll(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{

	})
}

