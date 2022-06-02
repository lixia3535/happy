package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"happy/models"
	"happy/utils"
	"log"
	"net/http"
	"time"
)

//自定义中间件第2种定义方式
func Authorized() gin.HandlerFunc{
	return func(c *gin.Context) {
		//验证权限 获取请求头里面的数据
		t := time.Now()
		// 查看url是否为adds,如若为adds则不需要验证authentication
		url := c.Request.URL.Path
		fmt.Println("url",url)
		if url == "/happy/adds"{
			c.Next()
		} else {
			//请求之前
			authentication := c.GetHeader("authentication")
			if authentication == "" {
				fmt.Println("权限认证失败")
				message := "权限认证失败"
				c.JSON(http.StatusUnauthorized,utils.GetResModel(http.StatusOK,message,nil))
				c.Abort()
				// return
			} else {
				c.Next()
				fmt.Println("我是自定义中间件第2种定义方式---请求之后")
				//请求之后
				//计算整个请求过程耗时
				t2 := time.Since(t)
				log.Println(t2)
			}

		}

	}
}
func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户的密码，然后base64解密后加密，
		user := new(models.User)
		c.ShouldBind(&user)
		pwd := utils.DecryptBase64(user.Password)
		user.Password = pwd

		fmt.Println(pwd)
		c.Next()
	}

}