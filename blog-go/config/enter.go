package config

type Config struct {
	Captcha Captcha `json:"captcha" yaml:"captcha"`
	Email   Email   `json:"email" yaml:"email"`
	ES      ES      `json:"es" yaml:"es"`
	Google  Google  `json:"gaode" yaml:"gaode"`
	Jwt     Jwt     `json:"jwt" yaml:"jwt"`
	Mysql   Mysql   `json:"mysql" yaml:"mysql"`
	SMMS    SMMS    `json:"smms" yaml:"smms"`
	Github  GitHub  `json:"qq" yaml:"qq"`
	Redis   Redis   `json:"redis" yaml:"redis"`
	System  System  `json:"system" yaml:"system"`
	Upload  Upload  `json:"upload" yaml:"upload"`
	Website Website `json:"website" yaml:"website"`
	Zap     Zap     `json:"zap" yaml:"zap"`
}
