package utils

import (
	"blog-go/global"
	"blog-go/model/appTypes"
	"blog-go/model/request"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

const (
	headerRefreshToken = "x-refresh-token"
	headerAccessToken  = "x-access-token"
)

func SetRefreshToken(c *gin.Context, token string, maxAge int) {
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	setCookie(c, headerRefreshToken, token, maxAge, host)
}

func ClearRefreshToken(c *gin.Context) {
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	setCookie(c, headerRefreshToken, "", -1, host)
}

func setCookie(c *gin.Context, name, value string, maxAge int, host string) {
	if net.ParseIP(host) != nil {
		c.SetCookie(name, value, maxAge, "/", "", false, true)
	} else {
		c.SetCookie(name, value, maxAge, "/", host, false, true)
	}
}

func GetAccessToken(c *gin.Context) string {
	token := c.Request.Header.Get(headerAccessToken)
	return token
}

func GetRefreshToken(c *gin.Context) string {
	token, _ := c.Cookie(headerRefreshToken)
	return token
}

func GetClaims(c *gin.Context) (*request.JwtCustomClaims, error) {
	token := GetAccessToken(c)
	j := NewJWT()
	claims, err := j.ParseAccessToken(token)
	if err != nil {
		global.Log.Error("Failed to retrieve JWT parsing information from Gin's Context. Please check if the request header contains 'x-access-token' and if the claims structure is correct.", zap.Error(err))
		return nil, err
	}
	return claims, nil
}

func GetRefreshClaims(c *gin.Context) (*request.JwtCustomRefreshClaims, error) {
	token := GetRefreshToken(c)
	j := NewJWT()
	claims, err := j.ParseRefreshToken(token)
	if err != nil {
		global.Log.Error("Failed to retrieve JWT parsing information from Gin's Context. Please check if the request header contains 'x-refresh-token' and if the claims structure is correct.", zap.Error(err))
		return nil, err
	}
	return claims, nil
}

func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.UserID
		}
	} else {
		waitUse := claims.(*request.JwtCustomClaims)
		return waitUse.UserID
	}
}

func GetUserInfo(c *gin.Context) *request.JwtCustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*request.JwtCustomClaims)
		return waitUse
	}
}

func GetRoleID(c *gin.Context) appTypes.RoleID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.RoleID
		}
	} else {
		waitUse := claims.(*request.JwtCustomClaims)
		return waitUse.RoleID
	}
}

func GetUUID(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*request.JwtCustomClaims)
		return waitUse.UUID
	}
}
