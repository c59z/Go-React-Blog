package router

import (
	"blog-go/api"
	"blog-go/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(privateRouter *gin.RouterGroup, publicRouter *gin.RouterGroup, adminRouter *gin.RouterGroup) {
	userRouter := privateRouter.Group("user")

	userPublicRouter := publicRouter.Group("user")
	userLoginRouter := publicRouter.Group("user").Use(middleware.LoginRecord())
	userAdminRouter := adminRouter.Group("user")

	userApi := api.ApiGroupApp.UserApi
	{
		userRouter.POST("logout", userApi.Logout)
		userRouter.PUT("resetPassword", userApi.UserResetPassword)
		userRouter.GET("info", userApi.UserInfo)
		userRouter.PUT("changeInfo", userApi.UserChangeInfo)
		userRouter.GET("weather", userApi.UserWeather)
		userRouter.GET("chart", userApi.UserChart)
	}
	{
		userPublicRouter.POST("forgotPassword", userApi.ForgotPassword)
		userPublicRouter.GET("card", userApi.UserCard)
	}
	{
		userLoginRouter.POST("register", userApi.Register)
		userLoginRouter.POST("login", userApi.Login)
	}
	{
		userAdminRouter.GET("list", userApi.UserList)
		userAdminRouter.PUT("freeze", userApi.UserFreeze)
		userAdminRouter.PUT("unfreeze", userApi.UserUnfreeze)
		userAdminRouter.GET("loginList", userApi.UserLoginList)
	}
}
