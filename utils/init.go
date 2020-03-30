package utils

import (
	"github.com/SergeyDavidenko/subscription/config"
	"github.com/SergeyDavidenko/subscription/db"
	log "github.com/sirupsen/logrus"
)

// InitServer ...
func InitServer(configPath string) {
	var err error
	config.Conf, err = config.LoadConf(configPath)
	if err != nil {
		log.Error(err)
	}
	level, err := log.ParseLevel(config.Conf.Log.Level)
	if err != nil {
		log.Error("Cannot parse log level")
		log.SetLevel(log.InfoLevel)
		config.LogLevel = log.InfoLevel
	}
	if level == log.DebugLevel {
		log.SetReportCaller(true)
	}
	log.Debug("Set log level ", level)
	config.LogLevel = level
	log.SetLevel(level)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: config.Conf.Log.DisableColors,
		FullTimestamp: config.Conf.Log.FullTimestamp,
	})

	db.ConnectorPoll()
	if config.Conf.Redis.UseRedis || config.Conf.API.UserRedisCache {
		if config.Conf.Redis.Address == "" {
			log.Warn("Redis address host not set use default value localhost")
			config.Conf.Redis.Address = "localhost"
		}
		if config.Conf.Redis.Port == 0 {
			config.Conf.Redis.Port = 6379
			log.Warn("Redis port not set use default port 6379")
		}
		db.RedisInitialize()
	}
}
