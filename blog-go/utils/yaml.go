package utils

import (
	"blog-go/global"
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"
)

var configFilePath = "./conf.yaml"

func LoadYAML() ([]byte, error) {
	return os.ReadFile(configFilePath)
}

func SaveYAML() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	return os.WriteFile(configFilePath, byteData, fs.ModePerm)
}
