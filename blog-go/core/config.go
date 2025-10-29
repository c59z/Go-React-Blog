package core

import (
	"blog-go/config"
	"blog-go/utils"
	"log"

	"gopkg.in/yaml.v3"
)

func InitConfig() *config.Config {
	c := &config.Config{}
	yamlBytes, err := utils.LoadYAML()
	if err != nil {
		log.Fatalf("failed to load yaml config: %v", err)
	}
	if err := yaml.Unmarshal(yamlBytes, &c); err != nil {
		log.Fatalf("failed to unmarshal yaml config: %v", err)
	}
	return c
}
