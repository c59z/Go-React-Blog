package database

import "blog-go/global"

type Login struct {
	global.MODEL
	UserID      uint   `json:"user_id"`
	User        User   `json:"user" gorm:"foreignKey:UserID"`
	LoginMethod string `json:"login_method"`
	IP          string `json:"ip"`
	Address     string `json:"address"`
	OS          string `json:"os"`
	DeviceInfo  string `json:"device_info"`
	BrowserInfo string `json:"browser_info"`
	Status      int    `json:"status"`
}
