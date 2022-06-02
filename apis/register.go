package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"happy/databases"
	"happy/models"
	"happy/utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

func LoginEndpoinit(c *gin.Context) {
	user := new(models.User)
	err := c.ShouldBind(&user)
	if err != nil{
		log.Fatal("登陆失败",err.Error())
	}
	userdb := databases.FindUserByUserName(user.Username,user.Password)
	if userdb.Id == 0 {
		c.JSON(http.StatusOK,utils.GetResModel(http.StatusUnauthorized,"用户名或者密码错误！",nil))
		fmt.Println("用户名或者密码错误！")
		return
	}

	c.JSON(http.StatusOK,utils.GetResModel(http.StatusOK,"登陆成功！",nil))

}

func AddUser(c *gin.Context){
	var message string
	// 获取前端请求的参数
	user := new(models.User)
	c.ShouldBind(&user)
	_,err := databases.AddUser(*user)
	if err != nil {
		c.JSON(http.StatusOK,utils.GetResModel(http.StatusOK,"用户名或者密码错误！",nil))
		fmt.Println("新增用户失败！",err.Error())
		message = "新增用户失败！"
	}
	message = "新增用户成功！"
	// log.Fatal("添加用户成功！")
	c.JSON(http.StatusOK,utils.GetResModel(http.StatusOK,message,nil))

}

func FindAllUser(c *gin.Context){
	users,err := databases.FindAllUser()
	if err != nil{
		c.JSON(http.StatusOK,utils.GetResModel(http.StatusOK,"查询所有用户失败！",nil))
		fmt.Println("查询所有用户失败！",err.Error())
	}
	c.JSON(http.StatusOK,utils.GetResModel(http.StatusOK,"查询成功",users))
}

func FindUserById(c *gin.Context){
	id := c.Query("id")
	idint,err := strconv.Atoi(id)
	if err != nil{
		fmt.Println("字符串转int失败",err.Error())
	}
	user,err := databases.FindUserById(idint)
	if err != nil {
		fmt.Println("查询用户失败",err.Error())
		c.JSON(http.StatusOK,utils.GetResModel(http.StatusInternalServerError,"查询用户失败",user))
		return
	}
	c.JSON(http.StatusOK,utils.GetResModel(http.StatusOK,"查询成功",user))
}

// 上传用户大的头像
func UploadFile(c *gin.Context){
	time.Sleep(20*time.Second)
	fmt.Println("sleep 2s")
	file,err := c.FormFile("upload")
	if err != nil{
		fmt.Println("获取上传的文件失败", err.Error())
		c.JSON(http.StatusOK,utils.GetResModel(http.StatusInternalServerError,"获取上传文件失败！",nil))
		return
	}
	// 获取文件的名字
	fileName := "./upload/" + file.Filename
	fmt.Println("文件的名字是：",fileName)
	// 将文件上传到服务器
	err = c.SaveUploadedFile(file,fileName)
	if err != nil {
		// 上传文件失败
		fmt.Println("文件上传失败",err.Error())
		c.JSON(http.StatusOK,utils.GetResModel(http.StatusInternalServerError,"获取上传文件失败！",nil))
		return
	}

	// 文件上传成功
	c.JSON(http.StatusOK,utils.GetResModel(http.StatusOK,"文件上传成功！！",fileName))
}

func UpdateUser(c *gin.Context){
	// 获取用户信息
	var user models.User
	c.BindJSON(&user)

	// 1、查询id对应的用户存不存在
	id := user.Id
	_, err := databases.FindUserById(id)
	if err != nil{
		log.Println("用户不存在，更新用户失败")
		c.JSON(http.StatusOK,utils.GetResModel(http.StatusInternalServerError,"用户不存在，更新用户失败！！",nil))
		return
	}
	// 添加更新时间
	nowData := time.Now()
	user.Updatetime=nowData
	info,err := databases.EditUser(user)
	if err != nil{
		c.JSON(http.StatusOK,utils.GetResModel(http.StatusInternalServerError,"更新用户失败！！"+err.Error(),nil))
		return
	}

	c.JSON(http.StatusOK,utils.GetResModel(http.StatusOK,"更新用户成功！！",info))

}

func DeleteUser(c *gin.Context) {
	var err error
	var conId int
	// 获取前端请求数据
	id, exist := c.GetQuery("id")
	if !exist {
		log.Println("获取请求数据失败")
	}

	conId,err = strconv.Atoi(id)
	if err != nil{
		fmt.Println(err.Error())
		c.JSON(http.StatusOK,utils.GetResModel(http.StatusInternalServerError,"删除用户失败！！", nil))
		return
	}
	_,err = databases.DeleteUser(conId)
	if err != nil{
		fmt.Println(err.Error())
		c.JSON(http.StatusOK,utils.GetResModel(http.StatusInternalServerError,"删除用户失败！！", nil))
		return
	}
	c.JSON(http.StatusOK,utils.GetResModel(http.StatusOK,"删除用户成功！！", nil))
}



