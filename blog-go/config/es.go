package config

type ES struct {
	URL            string `json:"url" yaml:"url"`                           // es url
	Username       string `json:"username" yaml:"username"`                 // es usename
	Password       string `json:"password" yaml:"password"`                 // es password
	IsConsolePrint bool   `json:"is_console_print" yaml:"is_console_print"` // print the Elasticsearch statement in the console
}
