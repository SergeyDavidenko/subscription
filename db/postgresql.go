package db

import (
	"context"
	"os"

	config "github.com/SergeyDavidenko/subscription/config"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

var (
	// DB pool connect
	DB  *pgxpool.Pool
	err error
)

// ConnectorConfig ...
func ConnectorConfig() *pgxpool.Config {
	logLevel, err := log.ParseLevel(config.Conf.PostgreSQL.LogLevel)
	if err != nil {
		log.Error("Cannot parse log level for postgresql, use default value warning")
		logLevel = log.WarnLevel

	}
	log.Debug("Log level: ", logLevel)
	logrusLogger := log.New()
	logrusLogger.SetLevel(logLevel)
	logrusLogger.WithFields(log.Fields{
		"module": "pgx",
	}).Info("Apend log field")
	logger := logrusadapter.NewLogger(logrusLogger)
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Error("Unable to parse DATABASE_URL", "error", err)
	}
	if poolConfig.ConnConfig.Database == "" {
		poolConfig.ConnConfig.Logger = logger
		poolConfig.ConnConfig.Host = config.Conf.PostgreSQL.Address
		poolConfig.ConnConfig.Port = uint16(config.Conf.PostgreSQL.Port)
		poolConfig.ConnConfig.Database = config.Conf.PostgreSQL.Database
		poolConfig.ConnConfig.User = config.Conf.PostgreSQL.Username
		poolConfig.ConnConfig.Password = config.Conf.PostgreSQL.Password
		poolConfig.ConnConfig.TLSConfig = nil
		poolConfig.MaxConns = config.Conf.PostgreSQL.MaxConnections
	}
	return poolConfig
}

// ConnectorPoll ...
func ConnectorPoll() {
	configConnect := ConnectorConfig()
	DB, err = pgxpool.ConnectConfig(context.Background(), configConnect)
	if err != nil {
		log.Error("Unable to create connection pool", "error", err)
		os.Exit(1)
	}
}
