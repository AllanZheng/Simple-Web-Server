package main
import(

	"github.com/gin-gonic/gin"
	"github/SimpleWeb/service"
	"flag"
	"fmt"
	"net/http"
)
func main(){
	
	flag.Parse()//
	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	router := ginReady()   //获得路由实例
	router.Run(":8080")

}
func ginReady() *gin.Engine{
    router := gin.New()
	//router.use(middleware.Auth())
	router.LoadHTMLGlob("templates/*")
    router.GET("/login", service.LoginForm)
    router.POST("/login", service.PostLoginForm)
	router.GET("/logout", service.Logout)
	r := router.Group("/service")
	r.Use(Auth())
	{
		r.GET("/", service.HomePage)
		r.GET("/createdep",service.CreateDep)
		r.GET("/updatedep",service.UpdateDep)
		r.GET("/deletedep",service.DeleteDep)
		r.GET("/searchdep",service.SearchDep)
		r.GET("/searchall",service.SearchAll)
	/*
		router.POST("/createuser",CreateUser)
		router.POST("/updateuser",updateUser)
		router.POST("/deleteuser",deleteUser)
		router.POST("/searchdep",searchUser)
	*/
	}

	fmt.Println("Server Start!")
	return router
}
/*
var sercets = gin.H{
	"admin": gin.H{"email":"amdin@example.com"},
}
*/
func Auth() gin.HandlerFunc{
	return func(c *gin.Context){
	username := c.Query("username")
	password := c.Query("password")
	fmt.Println("Middleware Using!")
	/*
	user := c.MustGet(gin.AuthUserKey).(string)
	authorized := r.Group("/",gin.BasicAuth(gin.Accounts{
		"admin":"12345",
	}))
	*/
	if username == "admin" && password =="12345" {
        c.JSON(http.StatusOK,gin.H{
            "message" : "Authorized!",
		})
		c.Next()
	}else{
	
		c.Abort()
	    c.HTML(
		http.StatusUnauthorized,"fail.html",gin.H{
			"message": "Authorized Failed!",
		})

	}
}

}

