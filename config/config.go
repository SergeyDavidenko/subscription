package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var defaultConf = []byte(`
core:
  environment: "dev"
  lable: "auth-service"
api:
  port: ":8088"
  health_port: ":1488"
  metric_uri: "/metrics"
  health_uri: "/healthz"
log:
  level: "info"
key:
  pub: "keys/app.rsa.pub"
  private: "keys/app.rsa"
`)

var (
	// Conf ...
	Conf ConfYaml
	// LogLevel ...
	LogLevel log.Level
)

// ConfYaml is config structure.
type ConfYaml struct {
	Core       SectionCore       `yaml:"core"`
	API        SectionAPI        `yaml:"api"`
	Log        SectionLog        `yaml:"log"`
	Redis      SectionRedis      `yaml:"redis"`
	PostgreSQL SectionPostgreSQL `yaml:"postgresql"`
	Auth       SectionAuth       `yaml:"auth"`
}

// SectionCore is sub section of config.
type SectionCore struct {
	Environment string `yaml:"environment"`
	Lable       string `yaml:"lable"`
}

// SectionAPI is sub section of config.
type SectionAPI struct {
	MetricURI          string `yaml:"metric_uri"`
	HealthURI          string `yaml:"health_uri"`
	UseAuth            bool   `yaml:"use_auth"`
	AuthLogin          string `yaml:"auth_login"`
	AuthPassword       string `yaml:"auth_password"`
	Port               int    `yaml:"port"`
	HealthPort         int    `yaml:"health_port"`
	PProfPort          int    `yaml:"pprof_port"`
	CokiesName         string `yaml:"cokies_name"`
	CokiesDomain       string `yaml:"cokies_domain"`
	TokenExpireMinutes int64  `yaml:"token_expire_minutes"`
	UserRedisCache     bool   `yaml:"user_redis_cache"`
}

// SectionLog is sub section of config.
type SectionLog struct {
	Level         string `yaml:"level"`
	DisableColors bool   `yaml:"disable_colors"`
	FullTimestamp bool   `yaml:"full_timestamp"`
}

// SectionRedis is sub section of config.
type SectionRedis struct {
	UseRedis bool   `yaml:"use_redis"`
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
}

// SectionPostgreSQL is sub section of config.
type SectionPostgreSQL struct {
	Address        string `yaml:"address"`
	Port           int    `yaml:"port"`
	Database       string `yaml:"database"`
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
	MaxConnections int32  `yaml:"max_connections"`
	LogLevel       string `yaml:"log_level"`
}

// SectionAuth is sub section of config.
type SectionAuth struct {
	URL string `yaml:"url"`
}

// LoadConf load config from file and read in environment variables that match
func LoadConf(confPath string) (ConfYaml, error) {
	var conf ConfYaml

	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix("go")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if confPath != "" {
		content, err := ioutil.ReadFile(confPath)

		if err != nil {
			return conf, err
		}

		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return conf, err
		}
	} else {
		// Search config in home directory with name ".gorush" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigName("config")

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			if err := viper.ReadConfig(bytes.NewBuffer(defaultConf)); err != nil {
				return conf, err
			}
		}
	}

	// Core
	conf.Core.Environment = viper.GetString("core.environment")
	conf.Core.Lable = viper.GetString("core.lable")

	//API
	conf.API.Port = viper.GetInt("api.port")
	conf.API.HealthPort = viper.GetInt("api.health_port")
	conf.API.PProfPort = viper.GetInt("api.pprof_port")
	conf.API.HealthURI = viper.GetString("api.health_uri")
	conf.API.MetricURI = viper.GetString("api.metric_uri")
	conf.API.UseAuth = viper.GetBool("api.use_auth")
	conf.API.AuthLogin = viper.GetString("api.auth_login")
	conf.API.AuthPassword = viper.GetString("api.auth_password")
	conf.API.CokiesName = viper.GetString("api.cokies_name")
	conf.API.CokiesDomain = viper.GetString("api.cokies_domain")
	conf.API.TokenExpireMinutes = viper.GetInt64("api.token_expire_minutes")
	conf.API.UserRedisCache = viper.GetBool("api.user_redis_cache")

	//Log
	conf.Log.Level = viper.GetString("log.level")
	conf.Log.DisableColors = viper.GetBool("log.disable_colors")
	conf.Log.FullTimestamp = viper.GetBool("log.full_timestamp")

	//Redis
	conf.Redis.Address = viper.GetString("redis.address")
	conf.Redis.Port = viper.GetInt("redis.port")
	conf.Redis.UseRedis = viper.GetBool("redis.use_redis")

	//Key
	conf.Auth.URL = viper.GetString("auth.url")

	//PostgreSQL
	conf.PostgreSQL.Address = viper.GetString("postgresql.address")
	conf.PostgreSQL.Port = viper.GetInt("postgresql.port")
	conf.PostgreSQL.Database = viper.GetString("postgresql.database")
	conf.PostgreSQL.Username = viper.GetString("postgresql.username")
	conf.PostgreSQL.Password = viper.GetString("postgresql.password")
	conf.PostgreSQL.MaxConnections = viper.GetInt32("postgresql.max_connections")
	conf.PostgreSQL.LogLevel = viper.GetString("postgresql.log_level")

	return conf, nil
}
