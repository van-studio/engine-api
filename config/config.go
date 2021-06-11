package config

type Config struct {
	Listen   string   `yaml:"listen"`
	Database Database `yaml:"database"`
}

type Database struct {
	Uri string `yaml:"uri"`
	Db  string `yaml:"db"`
}
