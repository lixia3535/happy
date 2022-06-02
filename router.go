package main

import (
	"github.com/gin-gonic/gin"
	"happy/apis"
)


func InitGin() *gin.Engine{
	//r := gin.Default()
	// r.GET("/user",getUser)
	// r.POST("/addUser",addUser)

	router :=gin.New()
	// 全局中间件，将日志写到gin.DefaultWriter,即使你设置了GIN_MODE=release
	router.Use(gin.Logger())
	// recovery中间件从任何panic恢复，如果出现panic。他会写一个500错误
	router.Use(gin.Recovery())
	router.Static("/upload", "./upload/")
	// router.Use(middleware.DbClient())

	// 每个路由的中间件，你能添加任意数量的中间键
	// router.GET("/benmark")

	// 这个是一个组
	authorized  := router.Group("/happy")
	// 给组内每一个路由添加中间键，类似于java的拦截器
	/*authorized .Use(middleware.Authorized())
	{*/
		authorized.POST("/adds",apis.LoginEndpoinit)
		// authorized.POST("/add",middleware.LoginHandler(),apis.AddUser)
		authorized.POST("/add",apis.AddUser)
		authorized.GET("/findAllUser",apis.FindAllUser)
		authorized.GET("/findUserById",apis.FindUserById)
		authorized.POST("/uploadFile",apis.UploadFile)
		authorized.POST("/updateUser",apis.UpdateUser)
		authorized.POST("/deleteUser",apis.DeleteUser)
		authorized.POST("/saveDictionaries",apis.SaveDictionaries)
		goods := authorized.Group("/goods")
		goods.GET("/goods")
	//}

	// router.Group("/happy",)
	return router
}


