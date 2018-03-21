package model

import "time"

//辩论话题
type Argue struct {
	ID 			uint 		`gorm:"primary_key" json:"id"`
	Name 		string 		`gorm:"size:128" json:"name"`
	Describe 	string 		`gorm:"size:1024" json:"describe"`    //话题背景描述
	IsHot 		int 			`json:"is_hot"`
	Followers 	string		`json:"followers"`
	CategoryId 	uint 		`json:"category_id"`    //话题所属的分类ID
	Categories 	[]Category 	`gorm:"many2many:argue_category;ForeignKey:ID;AssociationForeignKey:ID" json:"categories"` //话题可以对应多个分类
	Talk 		[]Talk 						  //话题对应多个评论
	UserID 		uint 		`json:"user_id"`   //创建话题的user
	User 		User 		`json:"user"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	*time.Time
}
