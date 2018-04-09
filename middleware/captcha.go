package middleware

import (
	"net/http"
	//"github.com/gin-gonic/gin"
	"github.com/lifei6671/gocaptcha"
	"fmt"
)
const(
	dx = 240
	dy= 120
)

func GetCaptcha(w http.ResponseWriter,r *http.Request){
	captchaImage,err := gocaptcha.NewCaptchaImage(dx,dy,gocaptcha.RandLightColor())
	captchaImage.DrawNoise(gocaptcha.CaptchaComplexLower)
	captchaImage.DrawText(gocaptcha.RandText(4))
	captchaImage.Drawline(2)
	captchaImage.DrawBorder(gocaptcha.ColorToRGB(0x17A7A7A))
	if err !=nil{
		fmt.Println(err)
	}
	captchaImage.SaveImage(w,gocaptcha.ImageFormatJpeg)
}