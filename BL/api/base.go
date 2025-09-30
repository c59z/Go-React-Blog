package api

import (
	"blog-go/global"
	"blog-go/model/request"
	"blog-go/model/response"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type BaseApi struct {
}

var store = base64Captcha.DefaultMemStore

func (b *BaseApi) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(
		global.Config.Captcha.Height,
		global.Config.Captcha.Width,
		global.Config.Captcha.Length,
		global.Config.Captcha.MaxSkew,
		global.Config.Captcha.DotCount,
	)

	captcha := base64Captcha.NewCaptcha(driver, store)

	id, b64s, err := captcha.Generate()

	if err != nil {
		global.Log.Error("Failed to generate captcha:", zap.Error(err))
		response.FailWithMessage("Failed to generate captcha", c)
		return
	}
	response.OkWithData(response.Captcha{
		CaptchaID: id,
		PicPath:   b64s,
	}, c)

}

func (b *BaseApi) SendEmailVerifcationCode(c *gin.Context) {
	var req request.SendEmailVerificationCode
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// if store.Verify(req.CaptchaID, req.Captcha, true) {
	// 	err := baseService.SendEmailVerifcationCode(c, req.Email)
	// 	if err != nil {
	// 		global.Log.Error("Failed to send email:", zap.Error(err))
	// 		response.FailWithMessage("Failed to send email", c)
	// 		return
	// 	}
	// 	response.OkWithMessage("Successfully sent email", c)
	// 	return
	// }
	// todo test
	err = baseService.SendEmailVerifcationCode(c, req.Email)
	if err != nil {
		global.Log.Error("Failed to send email:", zap.Error(err))
		response.FailWithMessage("Failed to send email", c)
		return
	}
	response.OkWithMessage("Successfully sent email", c)
}
