package config

//昵称格式    中英数字
const RegexpNickName = `^[a-zA-Z0-9\.\p{Han}]{1,50}$`

//手机格式
const RegexpPhone = `^1[0-9]{10}$`

//邮箱格式
const RegexpEmail = `^(\w)+(\.\w+)*@(\w)+((\.\w+)+)$`


//验证码session
const CaptchaSessionName = "__captcha__"
