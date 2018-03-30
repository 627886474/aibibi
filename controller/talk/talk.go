package talk

import (
	"github.com/gin-gonic/gin"
	"github.com/zl/aibibi/controller/common"
	"github.com/zl/aibibi/model"
	"unicode/utf8"
	"strconv"
	"net/http"
	"github.com/jinzhu/gorm"
)

//对话题添加talk
func Save(c *gin.Context){
	SendErrJSON := common.SendErrJSON
	var talk model.Talk
	var argue model.Argue

	type TalkReqData struct {
		Content string `json:"content" binding:"required"`
		Rank 	int 	 `json:"rank" binding:"required"`
		ArgueID uint 		`json:"argue_id" binding:"required"`
	}
	var talkData TalkReqData

	iuser,_ :=c.Get("user")
	user := iuser.(model.User)

	if err := c.ShouldBindJSON(&talkData);err !=nil{
		SendErrJSON("参数无效",c)
		return
	}

	if err := model.DB.First(&argue,talkData.ArgueID).Error;err !=nil{
		SendErrJSON("无效的话题",c)
		return
	}
	if talkData.Content == ""{
		SendErrJSON("回复不能为空",c)
		return
	}
	if utf8.RuneCountInString(talkData.Content) > model.MaxContentLen{
		msg := "回复不能超过"+strconv.Itoa(model.MaxContentLen)+"字符"
		SendErrJSON(msg,c)
		return
	}

	if talkData.Rank != model.Affirmative && talkData.Rank != model.Negative{
		SendErrJSON("请选择正方或者反方",c)
		return
	}

	talk.UserID = user.ID
	talk.Content = talkData.Content
	talk.Rank = talkData.Rank
	talk.ArgueID = talkData.ArgueID

	//创建talk
	if err := model.DB.Create(&talk).Error;err !=nil{
		SendErrJSON("error",c)
		return
	}

	//关联talk的argues 否则返回的json中argues为空
	if err := model.DB.Model(&talk).Related(&talk.Argues).Error;err !=nil{
		SendErrJSON("关联话题失败",c)
		return
	}

	//关联talk的user 否则返回的json中user为空
	if err := model.DB.Model(&talk).Related(&talk.User).Error;err !=nil{
		SendErrJSON("关联用户失败",c)
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"errNo":model.ErrorCode.SUCCESS,
		"msg":"success",
		"data":talk,
	})
}

//点赞talk
func UP(c *gin.Context) {
	SendErrJSON := common.SendErrJSON

	type UpData struct {
		ID        uint `json:"id" binding:"required"`
		ApplaudUp bool `json:"applaud_up"`
	}
	var upData UpData
	var talk model.Talk

	if err := c.ShouldBindJSON(&upData); err != nil {
		SendErrJSON("参数无效", c)
		return
	}
	if upData.ApplaudUp == true {
		if err := model.DB.Model(&talk).Where("id = ?", upData.ID).UpdateColumn("applaud_up", gorm.Expr("applaud_up + ?", 1)).Error; err != nil {
			SendErrJSON("点赞失败", c)
			return
		}
	} else {
		if err := model.DB.Model(&talk).Where("id = ?", upData.ID).UpdateColumn("applaud_down", gorm.Expr("applaud_down + ?", 1)).Error; err != nil {
			SendErrJSON("差评失败", c)
			return
		}

	}
	talk.ID = upData.ID
	if err := model.DB.Debug().Find(&talk).Where("id = ?",talk.ID).Error;err !=nil{
		SendErrJSON("error",c)
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"errNo":model.ErrorCode.SUCCESS,
		"msg":"success",
		"data":gin.H{
			"up":talk.ApplaudUp,
			"down":talk.ApplaudDown,
		},
	})
}