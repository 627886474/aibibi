package model

import (
	"github.com/garyburd/redigo/redis"
	"github.com/zl/aibibi/config"
	"time"
	"github.com/jinzhu/gorm"
	"fmt"
	"log"
	"os"
)


// redis 连接池
var RedisPool *redis.Pool

func initRedis(){
	RedisPool = &redis.Pool{
		MaxIdle :			config.RedisConfig.MaxIdle,
		MaxActive : 			config.RedisConfig.MaxActive,
		IdleTimeout : 240 *time.Second,
		Dial:func()(redis.Conn,error){
			c,err :=redis.Dial("tcp",config.RedisConfig.URL)
			if err !=nil{
				return nil,err
			}
			return c,nil
		},
	}
}

// DB数据库连接
var DB *gorm.DB

func initDB(){
	db ,err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err !=nil{
		fmt.Println(err.Error())
		log.Fatal(err.Error())
		os.Exit(-1)
	}
	if config.ServerConfig.Env == ProductionMode {
		db.LogMode(true)
	}
	//设置数据库连接参数
	db.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)
	DB = db

	// 设置默认表头 sh_
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "sh_" + defaultTableName
	}
}

//注册表 添加model后，在这里需要添加数据结构体
func RegisterDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:" + config.DBConfig.Password +"@tcp("+ config.DBConfig.Host+")/go_project?charset=utf8&parseTime=true&loc=Asia%2FShanghai")

	if err == nil {
		DB = db
		db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
		// 添加自动注册的表
		db.AutoMigrate(&User{},&Category{},Argue{},Talk{})
		//修改表的结构
		//db.Model(&User{}).ModifyColumn("role","int")
		//db.Model(&Good{}).ModifyColumn("goods_id","int")

		return db, err
	}
	return nil, err
}


func init(){
	initRedis()
	initDB()
	RegisterDB()
}