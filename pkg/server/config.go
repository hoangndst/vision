package server

import "gorm.io/gorm"

type Config struct {
	DB          *gorm.DB
	Port        int
	LogFilePath string
	AutoMigrate bool
	GithubToken string
}

func NewConfig() *Config {
	return &Config{}
}
