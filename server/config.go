package server

import "gorm.io/gorm"

type Config struct {
	DB          *gorm.DB
	Port        int
	LogFilePath string
	AutoMigrate bool
}

func NewConfig() *Config {
	return &Config{}
}
