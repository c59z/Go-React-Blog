package config

// Redis cache database configuration
type Redis struct {
	Address  string `json:"address" yaml:"address"`   // Redis server address, usually "localhost:6379" or another host:port
	Password string `json:"password" yaml:"password"` // Password for Redis connection, leave empty if none
	DB       int    `json:"db" yaml:"db"`             // Database index to use, default is 0 in single-instance mode
}
