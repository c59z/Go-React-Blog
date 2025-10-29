package config

import (
	"strconv"
	"strings"

	"gorm.io/gorm/logger"
)

// Mysql database configuration
type Mysql struct {
	Host         string `json:"host" yaml:"host"`                     // Database server address
	Port         int    `json:"port" yaml:"port"`                     // Database server port
	Config       string `json:"config" yaml:"config"`                 // Database connection parameters, such as driver, charset, etc.
	DBName       string `json:"db_name" yaml:"db_name"`               // Name of the database to connect to
	Username     string `json:"username" yaml:"username"`             // Username for database connection
	Password     string `json:"password" yaml:"password"`             // Password for database connection
	MaxIdleConns int    `json:"max_idle_conns" yaml:"max_idle_conns"` // Maximum number of idle connections in the pool
	MaxOpenConns int    `json:"max_open_conns" yaml:"max_open_conns"` // Maximum number of open connections allowed
	LogMode      string `json:"log_mode" yaml:"log_mode"`             // Log mode, e.g., "info" or "silent", controls logging output
}

func (m Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DBName + "?" + m.Config
}
func (m Mysql) LogLevel() logger.LogLevel {
	switch strings.ToLower(m.LogMode) {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
}
