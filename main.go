package main
import(
	"fmt"
	"github.com/gin-gonic/gin"
	"github/SimpleWeb/entity"
	"github/SimpleWeb/util"
	"net/http"
	"flag"
)
func main(){

	flag.Parse()//
	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
    router := gin.Default()    //获得路由实例
    
    //添加中间件
    //router.Use(Middleware)
    //注册接口
    
	router.LoadHTMLGlob("templates/*")
	//测试

	router.GET("/login",func(c *gin.Context){
		c.HTML(http.StatusOK,"index.html",gin.H{
		})
	})
	router.GET("/update",func(c *gin.Context){
		c.HTML(http.StatusOK,"update.html",gin.H{
		})
	})
	router.GET("/search",func(c *gin.Context){
		c.HTML(http.StatusOK,"search.html",gin.H{
		})
	})
	router.GET("/delete",func(c *gin.Context){
		c.HTML(http.StatusOK,"delete.html",gin.H{
		})
	})
	router.POST("/login",func(c *gin.Context){
		var e entity.Info
		var err error
		e.Name = c.PostForm("Name")
		//e.ID = c.PostForm("ID")
		fmt.Println(e.Name)
	   
		if err = util.CreateForm(&e);err!=nil{
		c.HTML(http.StatusOK,"fail.html",gin.H{
		"Error Info!" : err,
		})
		}else{
			fmt.Println("Sucessful!")
			c.HTML(http.StatusOK,"res.html",gin.H{
		
				})
		}
		
	})
	router.POST("/delete",func(c *gin.Context){
		var e entity.Info
		var err error
		e.Name = c.PostForm("Name")
		//e.ID = c.PostForm("ID")
		fmt.Println(e.Name)

		if err = util.DeleteForm(&e);err!=nil{
		c.HTML(http.StatusOK,"fail.html",gin.H{
		"Error Info!" : err,
		})
		}else{
			c.HTML(http.StatusOK,"res.html",gin.H{})
			fmt.Println(" Delete Sucessful!")
		}
		
	})
	router.POST("/update",func(c *gin.Context){
		var e entity.Info
		var ne entity.Info
		var err error
		e.Name = c.PostForm("Name")
		ne.Name = c.PostForm("NewName")
		fmt.Println(e.Name)
        fmt.Println(ne.Name)
		if err = util.UpdateForm(&e,&ne);err!=nil{
		c.HTML(http.StatusOK,"fail.html",gin.H{
		"Error Info!" : err,
		})
		}else{
			c.HTML(http.StatusOK,"res.html",gin.H{})
			fmt.Println(" Update Sucessful!")
		}
		
	})
	router.POST("/search",func(c *gin.Context){
		var e entity.Info
		var err error
		e.Name = c.PostForm("Name")
		//e.ID = c.PostForm("ID")
		fmt.Println(e.Name)
        var res entity.Info
		if err,res = util.FindForm(&e);err!=nil{
		c.HTML(http.StatusOK,"fail.html",gin.H{
		"Error Info!" : err,
		})
		}else{
			c.HTML(http.StatusOK,"res.html",gin.H{
				"res": res,
			})
			fmt.Println(" Find Sucessful!")
		}
		
	})
	router.Run(":8080")
	fmt.Println("Server Start!")
}