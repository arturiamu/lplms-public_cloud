package config

import (
	"fmt"
	"github.com/spf13/viper"
	"net/url"
	"os"
	"path"
)

var AppConfig *Config

type Config struct {
	App     App     `yaml:"app" mapstructure:"app"`
	MySQL   MySQL   `yaml:"mysql" mapstructure:"mysql"`
	SQLLite SQLLite `yaml:"sqllite" mapstructure:"sqllite"`
}

type App struct {
	Env     string `yaml:"env" mapstructure:"env"`
	Name    string `yaml:"name" mapstructure:"name"`
	Version string `yaml:"version" mapstructure:"version"`
	Port    int    `yaml:"port" mapstructure:"port"`
}

type SQLLite struct {
	Path string `yaml:"path" mapstructure:"path"`
	Name string `yaml:"name" mapstructure:"name"`
}

type MySQL struct {
	Host            string `yaml:"host" mapstructure:"host"`
	Port            int    `yaml:"port" mapstructure:"port"`
	Username        string `yaml:"username" mapstructure:"username"`
	Password        string `yaml:"password" mapstructure:"password"`
	Database        string `yaml:"database" mapstructure:"database"`
	MaxIdleConn     int    `yaml:"max_idle_conn" mapstructure:"max_idle_conn"`
	MaxOpenConn     int    `yaml:"max_open_conn" mapstructure:"max_open_conn"`
	ConnMaxLifeTime int    `yaml:"conn_max_life_time" mapstructure:"conn_max_life_time"`
}

func LoadConfig(confPath string) (err error) {
	viper.SetConfigType("yaml")
	if confPath == "" {
		confPath, err = os.Getwd()
		if err != nil {
			return
		}
		confPath = path.Join(confPath, "/config/config.dev.yaml")
	}

	f, err := os.Open(confPath)
	if err != nil {
		return
	}
	defer f.Close()

	if err = viper.ReadConfig(f); err != nil {
		return
	}

	var cfg *Config
	if err = viper.Unmarshal(&cfg); err != nil {
		return
	}
	AppConfig = cfg
	return
}

func GetDSN() string {
	if AppConfig == nil {
		return ""
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=%s&parseTime=true",
		AppConfig.MySQL.Username, AppConfig.MySQL.Password, AppConfig.MySQL.Host,
		AppConfig.MySQL.Port, AppConfig.MySQL.Database, url.QueryEscape("Asia/Shanghai"))
}
