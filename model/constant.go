package model

const(
	// 最大名称长度
	MaxNameLen = 10
	//最大内容长度
	MaxContentLen = 5000
)


// redis相关常量，为了防止从redis中存取数据时key混乱,在此集中定义常量来作为各key的名字
const(
	//用户信息
	LoginUser = "loginUser"
	//resettime 重置密码的链接
	ResetTime = "resetTime"
	//ArgueMinuteLimit 每分钟最多能发表的话题数
	ArgueMinuteLimit = "argueMinuteLimit"
	//ArgueDayLimit 每分钟最多能发表的话题数
	ArgueDayLimit = "argueDayLimit"
	//TalkMinuteLimit 每分钟最多能发表的话题数
	TalkMinuteLimit = "talkMinuteLimit"
	//TalkDayLimit 每分钟最多能发表的话题数
	TalkDayLimit = "talkDayLimit"
)