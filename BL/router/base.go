package router

import (
	"blog-go/api"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (b *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("v1")
	baseApi := api.ApiGroupApp.BaseApi
	{
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.POST("sendEmailVerifcationCode", baseApi.SendEmailVerifcationCode)
		baseRouter.GET("githubLoginURL", baseApi.GithubLoginURL)
	}
}
