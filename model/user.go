package model

import (
	"time"
	"strconv"
	"fmt"
	"crypto/md5"
	"github.com/zl/aibibi/config"
	"github.com/json-iterator/go"
	"errors"
	"github.com/garyburd/redigo/redis"
)

type User struct{
	ID 				uint 		`gorm:"primary_key" json:"id"`
	Name 			string 		`gorm:"unique_index;size(64)" json:"name"`
	PassWord 		string 		`json:"-"`
	Birthday 		time.Time	`gorm:"default:null" json:"birthday"`
	Sex 				int 	 		`json:"sex"`
	Mobile 			string  		`gorm:"default:null" json:"mobile"`
	Email 			string 		`json:"email"`
	Role 			int 			`json:"role"`  //用户角色
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}


//每个用户一个唯一的salt
func (u User) Salt() string{
	var userSalt string
	if u.PassWord == ""{
		userSalt = strconv.Itoa(int(time.Now().Unix()))
	}else {
		userSalt = u.PassWord[0:10]
	}
	return userSalt
}

//密码加密
func (u User)EncryptPassword(password,salt string)(hash string){
	password = fmt.Sprintf("%x",md5.Sum([]byte(password)))
	hash = salt+password+config.ServerConfig.PassSalt
	hash = salt+fmt.Sprintf("%x",md5.Sum([]byte(hash)))
	return
}

//密码是否正确
func (u User) CheckPassword(password string) bool{
	if password =="" || u.PassWord ==""{
		return false
	}
	return u.EncryptPassword(password,u.Salt()) == u.PassWord
}



// 将用户信息存到redis
func UserToredis(user User)error{
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	userBytes,err := json.Marshal(user)

	if err !=nil{
		fmt.Println(err)
		return errors.New("error")
	}
	loginUserKey := fmt.Sprintf("%s%d",LoginUser,user.ID)
	// redis 连接
	RedisConn := RedisPool.Get()
	defer RedisConn.Close()

	if _,redisErr :=RedisConn.Do("SET",loginUserKey,userBytes,"Ex",config.ServerConfig.TokenMaxAge);redisErr !=nil{
		fmt.Println("redis set failed: ",redisErr.Error())
		return errors.New("error")
	}
	return nil
}

//从redis读取用户信息
func UserFromRedis(userID int)(User,error){
	LoginUser := fmt.Sprintf("%s%d",LoginUser,userID)

	RedisConn := RedisPool.Get()
	defer RedisConn.Close()

	userBytes,err := redis.Bytes(RedisConn.Do("GET",LoginUser))
	if err !=nil{
		return User{},errors.New("未登录")
	}

	var user User
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	bytesErr := json.Unmarshal(userBytes,&user)
	if bytesErr !=nil{
		fmt.Println(bytesErr)
		return user,errors.New("未登录")
	}
	return user,nil
}

const (
	// 普通用户
	UserRoleNormal = 1
	// 管理员
	UserRoleAdmin = 2
)

const (
	//男
	UserSexMale = 0
	//女
	UserSexFemale =1
	// 个性签名最大长度
	MaxSignatureLen = 200
)