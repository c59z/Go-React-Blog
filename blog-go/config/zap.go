package config

// Zap logger configuration
type Zap struct {
	Level          string `json:"level" yaml:"level"`                       // Log level, usually "info"
	Filename       string `json:"filename" yaml:"filename"`                 // Log file path
	MaxSize        int    `json:"max_size" yaml:"max_size"`                 // Max log file size (MB) before rotation
	MaxBackups     int    `json:"max_backups" yaml:"max_backups"`           // Max number of old log files to keep
	MaxAge         int    `json:"max_age" yaml:"max_age"`                   // Max days to keep old log files
	IsConsolePrint bool   `json:"is_console_print" yaml:"is_console_print"` // Print logs to console if true
}
