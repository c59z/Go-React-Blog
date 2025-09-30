package initialize

import (
	"blog-go/global"
	"blog-go/router"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()
	// TODO

	store := cookie.NewStore([]byte(global.Config.System.SessionsSecret))
	Router.Use(sessions.Sessions("session", store))

	routerGroup := router.RouterGroupApp
	publicGroup := Router.Group(global.Config.System.RouterPrefix)
	{
		routerGroup.InitBaseRouter(publicGroup)
	}
	return Router
}
