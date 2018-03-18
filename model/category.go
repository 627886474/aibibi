package model

import "time"

//话题的类别
type Category struct{
	ID 			uint 		`gorm:"primary_key" json:"id"`
	Name 		string 		`gorm:"unique_index" json:"name"`
	IsHot 		int			`json:"is_hot"`
	Followers 	string		`json:"followers"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

const (
	//热门话题
	HotYes = 0
	//冷门话题
	HotNo = 1
)