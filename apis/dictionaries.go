package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"happy/databases"
	"happy/models"
	"happy/utils"
	"net/http"
)

func SaveDictionaries(c *gin.Context){
	// 获取前台传过来的数据
	var dictionaries models.Dictionaries
	err := c.BindJSON(&dictionaries)
	// 解析json
	if err != nil{
		fmt.Println("json解析报错",err.Error())
		c.JSON(http.StatusOK,utils.GetResModel(http.StatusInternalServerError,err.Error(),nil))
	}
	err = databases.AddDictionaries(dictionaries)
	if err != nil {
		fmt.Println("数据插入数据库报错",err.Error())
		c.JSON(http.StatusOK,utils.GetResModel(http.StatusInternalServerError,err.Error(),nil))
	}
	c.JSON(http.StatusOK,utils.GetResModel(http.StatusOK,"数据插入成功！",nil))
}
