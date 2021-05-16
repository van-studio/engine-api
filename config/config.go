package config

import (
	"api/config/options"
	"github.com/kainonly/gin-extra/cors"
)

type Config struct {
	Listen   string                 `yaml:"listen"`
	Database options.DatabaseOption `yaml:"database"`
	Redis    options.RedisOption    `yaml:"redis"`
	Cors     cors.Option            `yaml:"cors"`
}
