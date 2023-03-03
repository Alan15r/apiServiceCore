package config

import (
	"encoding/json"
	"github.com/Alan15r/apiServiceCore/pkg/repository"
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	LogLevel string              `json:"loglevel"`
	Port     int                 `json:"port"`
	Database repository.Database `json:"database"`
}

func InitConfig() (*Config, error) {
	config := new(Config)
	if byteValue, err := os.ReadFile("configs/config.json"); err == nil {
		if err = json.Unmarshal(byteValue, config); err == nil {
			logrus.WithFields(logrus.Fields{
				"loglevel": config.LogLevel,
				"port":     config.Port,
				"db_host":  config.Database.Host,
				"db_port":  config.Database.Port,
				"db_name":  config.Database.Name,
				"db_user":  config.Database.UserName,
			}).Info("Config")
			return config, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}
