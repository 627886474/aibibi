package argue

import (
	"github.com/gin-gonic/gin"
	"github.com/zl/aibibi/controller/common"
	"github.com/zl/aibibi/model"
	"net/http"
	"fmt"
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
		fmt.Println(err.Error())
		SendErrJSON("参数无效",c)
		return
	}

	if argueData.Name == ""{
		SendErrJSON("话题名不能为空",c)
		return
	}
	if argueData.Describe == ""{
		SendErrJSON("描述不能为空",c)
		return
	}

	var category model.Category
	var argue model.Argue
	if err :=model.DB.First(&category,argueData.CategoryId).Error; err !=nil{
		SendErrJSON("无效的分类id",c)
		return
	}

	argue.Name = argueData.Name
	argue.Describe = argueData.Describe
	argue.CategoryId = argueData.CategoryId
	argue.Categories = argueData.Categories


	for i := 0; i < len(argue.Categories); i++ {
		var category model.Category
		if err := model.DB.First(&category, argue.Categories[i].ID).Error; err != nil {
			SendErrJSON("无效的版块id", c)
			return
		}
		argue.Categories[i] = category
	}

	if err := model.DB.Debug().Save(&argue).Error;err !=nil{
		SendErrJSON("error",c)
		return
	}
	//if err := model.DB.Debug().Model(&argue).Related(&category,"Categories");err !=nil{
	//	fmt.Println(err)
	//	SendErrJSON("error2222",c)
	//	return
	//}

	

	c.JSON(http.StatusOK,gin.H{
		"errorNo":model.ErrorCode.SUCCESS,
		"msg":"success",
		"data":gin.H{
			"argue":argue,
		},
	})
}

//创建话题
func Create(c *gin.Context){
	Save(c,false)
}
