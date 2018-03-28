package model

import "time"

//话题的辩论
type Talk struct {
	ID 			uint 		`gorm:"primary_key" json:"id"`
	Content 		string 		`json:"content"` 		//辩论内容
	Rank 		int 			`json:"rank"` 			//正方或者反方
	ApplaudUp 	int 			`gorm:"default:0"` 		//点赞数
	ApplaudDown 	int			`gorm:"default:0"` 	//差评数
	ArgueID 		uint 			`json:"argue_id"`		//辩论话题的ID
	Argues 		Argue				   				//辩论的内容对应一个话题
	UserID 		uint 								//不需要传userid,直接通过登录账号关联用户id
	User 		User
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	*time.Time
}


const(
	//正方
	Affirmative = 1
	//反方
	Negative = 2
)