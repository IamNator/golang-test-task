package config

import (
	"errors"
	"os"
	"sync"
)

type (
	Config struct {
		RedisURL    string `json:"-"`
		RabbitMQURL string `json:"-"`
		QueueName   string `json:"-"`
		PusherPORT  string `json:"-"`
		ReportPORT  string `json:"-"`
		WorkerPORT  string `json:"-"`
	}
)

var (
	configuration Config
	SyncOnce      sync.Once
)

func GetConfig() Config {
	return configuration
}

func InitConfig() error {
	var err error
	SyncOnce.Do(func() {
		if configuration.RedisURL, err = getEnv("REDIS_URL"); err != nil {
			return
		}
		if configuration.RabbitMQURL, err = getEnv("RABBITMQ_URL"); err != nil {
			return
		}
		if configuration.QueueName, err = getEnv("QUEUE_NAME"); err != nil {
			return
		}
		if configuration.PusherPORT, err = getEnv("PUSHER_PORT"); err != nil {
			return
		}
		if configuration.ReportPORT, err = getEnv("REPORT_PORT"); err != nil {
			return
		}
		if configuration.WorkerPORT, err = getEnv("WORKER_PORT"); err != nil {
			return
		}
	})

	return err
}

func getEnv(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return "", errors.New(key + " : is required")
	}

	return v, nil
}
