package category

import (
	"github.com/gin-gonic/gin"
	"github.com/zl/aibibi/controller/common"
	"github.com/zl/aibibi/model"
	"strings"
	"unicode/utf8"
	"strconv"
	"net/http"
	"github.com/microcosm-cc/bluemonday"
)

func Save(c *gin.Context,isEdit bool){
	SendErrJSON := common.SendErrJSON

	var category model.Category
	if err := c.ShouldBindJSON(&category);err !=nil{
		SendErrJSON("参数无效",c)
		return
	}
	category.Name = bluemonday.UGCPolicy().Sanitize(category.Name)   // 过滤掉不信任的内容
	category.Name = strings.TrimSpace(category.Name)

	if category.Name == ""{
		SendErrJSON("分类不能为空",c)
		return
	}
	if utf8.RuneCountInString(category.Name) > model.MaxNameLen{
		msg:="分类名称不能超过"+strconv.Itoa(model.MaxNameLen)+"个字符"
		SendErrJSON(msg,c)
		return
	}

	if err :=model.DB.Where("name = ?",category.Name).Find(&category).Error;err ==nil{
			SendErrJSON("已存在分类",c)
			return
	}


	var updateCateGory model.Category
	if !isEdit{
		//创建分类
		if err := model.DB.Create(&category).Error;err !=nil{
			SendErrJSON("error",c)
			return
		}
	}else {
		if err := model.DB.First(&updateCateGory,category.ID).Error;err ==nil{
			updateMap := make(map[string]interface{})
			updateMap["name"] = category.Name
			if err := model.DB.Model(&updateCateGory).Update(updateMap).Error;err !=nil{
				SendErrJSON("error",c)
				return
			}
		}else {
			SendErrJSON("更新失败",c)
			return
		}
	}

	var categoryJSON model.Category
	if isEdit{
		categoryJSON = updateCateGory
	}else {
		categoryJSON = category
	}
	c.JSON(http.StatusOK,gin.H{
		"errorNo":model.ErrorCode.SUCCESS,
		"msg":"success",
		"data":gin.H{
			"category":categoryJSON,
		},
	})
}

//创建分类
func Create(c *gin.Context){
	Save(c,false)
}

//更新分类
func Update(c *gin.Context){
	Save(c,true)
}

//获取一个分类的详情
func Info(c *gin.Context){
	SendErrJSON := common.SendErrJSON
	id,err := strconv.Atoi(c.Param("id"))
	if err !=nil{
		SendErrJSON("不存在的分类ID",c)
		return
	}

	var category model.Category
	if err := model.DB.First(&category,id).Error;err !=nil{
		SendErrJSON("不存在的分类ID",c)
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"errNo":model.ErrorCode.SUCCESS,
		"msg":"success",
		"data":gin.H{
			"category":category,
		},
	})
}

//获取所有分类
func List(c *gin.Context){
	SendErrJSON := common.SendErrJSON
	var categories []model.Category

	if err := model.DB.Order("followers asc").Find(&categories).Error;err !=nil{
		SendErrJSON("error",c)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"errNo":model.ErrorCode.SUCCESS,
		"msg":"success",
		"data":gin.H{
			"categories":categories,
		},
	})
}