package argue

import (
	"github.com/gin-gonic/gin"
	"github.com/zl/aibibi/controller/common"
	"github.com/zl/aibibi/model"
	"net/http"
	"fmt"
	"strconv"
)

func Save(c *gin.Context,isEdit bool){
	SendErrJSON := common.SendErrJSON

	type ArgueReqData struct{
		Name 	string 	`json:"name" binding:"required"`
		Describe string	`json:"describe" binding:"required"`
		CategoryId 	uint 	`json:"category_id" binding:"required"`
		Categories  []model.Category `json:"categories"`
	}

	var argueData ArgueReqData
	if err :=c.ShouldBindJSON(&argueData);err !=nil{
		fmt.Println(err)
		SendErrJSON("参数无效",c)
		return
	}

	/*
	请求体重加入了  required   标签，为必填，所以在请求格式中如果
	"name":""，"describe":"",  这种格式的数据直接会被判断为 "参数无效",
	所以下面的判空可以省略
	 */
	//if argueData.Name == ""{
	//	SendErrJSON("话题名不能为空",c)
	//	return
	//}
	//if argueData.Describe == ""{
	//	SendErrJSON("描述不能为空",c)
	//	return
	//}


	userInter,_ :=c.Get("user")
	user := userInter.(model.User)

	if err := model.DB.First(&user,user.ID).Error;err !=nil{
		SendErrJSON("error",c)
		return
	}

	var argue model.Argue
	var updateArgue model.Argue

	if isEdit ==false{
		argue.Name = argueData.Name
		argue.Describe = argueData.Describe
		argue.CategoryId = argueData.CategoryId
		argue.Categories = argueData.Categories
		argue.UserID = user.ID

		for i := 0; i < len(argue.Categories); i++ {
			var category model.Category
			if err := model.DB.First(&category, argue.Categories[i].ID).Error; err != nil {
				SendErrJSON("无效的分类id", c)
				return
			}
			argue.Categories[i] = category
		}

		if err := model.DB.Where("name = ?",argue.Name).Find(&argue).Error;err ==nil{
			SendErrJSON("已存在相同的讨论话题",c)
			return
		}
		if err := model.DB.Save(&argue).Error;err !=nil{
			fmt.Println(err)
			SendErrJSON("保存话题失败",c)
			return
		}
	}else {
		id,err := strconv.Atoi(c.Param("id"))
		if err !=nil{
			SendErrJSON("不存在的话题ID",c)
			return
		}
		if err := model.DB.First(&updateArgue,id).Error;err ==nil{
			updateArgue.Name = argueData.Name
			updateArgue.Describe = argueData.Describe
			updateArgue.CategoryId = argueData.CategoryId
			updateArgue.Categories = argueData.Categories

			if err := model.DB.Model(&argue).Where("id = ?",id).Update(map[string]interface{}{
				"name":updateArgue.Name,
				"describe":updateArgue.Describe,
				"category_id":updateArgue.CategoryId,
			}).Error;err !=nil{
				SendErrJSON("更新话题失败",c)
				return
			}

			var sql = "update  sh_argue_category set category_id = ? where argue_id = ?"
			if err := model.DB.Debug().Exec(sql,updateArgue.CategoryId,id).Error;err !=nil{
				SendErrJSON("更新话题分类失败",c)
			}

		}
	}

	c.JSON(http.StatusOK,gin.H{
		"errorNo":model.ErrorCode.SUCCESS,
		"msg":"success",
		"data":gin.H{
			"name":argue.Name,
			"describe":argue.Describe,

		},
	})
}

//创建话题
func Create(c *gin.Context){
	Save(c,false)
}

//更新话题
func Update(c *gin.Context){
	Save(c,true)
}