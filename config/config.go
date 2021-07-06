package config

import (
	"github.com/kainonly/gin-helper/authx"
	"github.com/kainonly/gin-helper/cookie"
)

type Config struct {
	Listen   string        `yaml:"listen"`
	Database Database      `yaml:"database"`
	Auth     authx.Option  `yaml:"auth"`
	Cookie   cookie.Option `yaml:"cookie"`
}

type Database struct {
	Dsn             string `yaml:"dsn"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
	TablePrefix     string `yaml:"table_prefix"`
}
