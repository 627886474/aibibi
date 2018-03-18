package model

import "time"

//话题的辩论
type Talk struct {
	ID 			uint 		`gorm:"primary_key" json:"id"`
	Content 		string 		`json:"content"` 		//辩论内容
	Rank 		int 			`json:"rank"` 			//正方或者反方
	Applaud 		int 			`json:"applaud"` 		//点赞或者差评
	ArgueID 		int 			`json:"argue_id"`		//辩论话题的ID
	Argues 		Argue				   				//辩论的内容对应一个话题
	UserID 		uint 		`json:"user_id"`
	User 		User
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	*time.Time
}

const (
	//赞同
	ApplaudUp = 0
	//差评
	ApplaudDown = 1
)

const(
	//正方
	Affirmative = 0
	//反方
	Negative = 1
)