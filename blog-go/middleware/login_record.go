package middleware

import (
	"blog-go/global"
	"blog-go/model/database"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ua-parser/uap-go/uaparser"
	"go.uber.org/zap"
)

// LoginRecord logs login activity as middleware.
func LoginRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Extract what we need BEFORE starting goroutine
		var userID uint
		if value, exists := c.Get("user_id"); exists {
			if id, ok := value.(uint); ok {
				userID = id
			}
		}

		ip := c.ClientIP()
		loginMethod := c.DefaultQuery("flag", "email")
		userAgent := c.Request.UserAgent()
		status := c.Writer.Status()

		// Run async logging safely (no closure over c)
		go func(userID uint, ip, loginMethod, userAgent string, status int) {
			address := getAddressFromIP(ip)
			os, deviceInfo, browserInfo := parseUserAgent(userAgent)

			login := database.Login{
				UserID:      userID,
				LoginMethod: loginMethod,
				IP:          ip,
				Address:     address,
				OS:          os,
				DeviceInfo:  deviceInfo,
				BrowserInfo: browserInfo,
				Status:      status,
			}

			if err := global.DB.Create(&login).Error; err != nil {
				global.Log.Error("Failed to record login", zap.Error(err))
			}
		}(userID, ip, loginMethod, userAgent, status)
	}
}

// getAddressFromIP returns the IP location.
// Currently returns a default value (no external API used).
func getAddressFromIP(ip string) string {
	fmt.Println("getAddressFromIP:" + ip)
	return "unknown"
}

// parseUserAgent extracts OS, device, and browser info from User-Agent.
func parseUserAgent(userAgent string) (os, deviceInfo, browserInfo string) {
	parser := uaparser.NewFromSaved()
	cli := parser.Parse(userAgent)
	os = cli.Os.Family
	deviceInfo = cli.Device.Family
	browserInfo = cli.UserAgent.Family
	return
}
