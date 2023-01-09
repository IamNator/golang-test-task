package report

import (
	"encoding/json"
	"sort"
	"twitch_chat_analysis/sdk/model"
	"twitch_chat_analysis/sdk/pkg/redis"
)

type (
	App struct {
		redisClient *redis.Redis
	}
)

func NewApp(redisClient *redis.Redis) *App {

	return &App{redisClient: redisClient}
}

func (r App) GetMessages() ([]model.Message, error) {
	// Connect to Redis

	// Get all keys in Redis
	keys, err := r.redisClient.Client.Keys("*").Result()
	if err != nil {
		return nil, err
	}

	// Get the values for all keys
	values, err := r.redisClient.Client.MGet(keys...).Result()
	if err != nil {
		return nil, err
	}

	// Parse the values into Message objects
	var messages []model.Message
	for _, value := range values {
		var m model.Message
		err := json.Unmarshal([]byte(value.(string)), &m)
		if err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}

	// Sort the messages in chronological descending order
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].Timestamp > messages[j].Timestamp
	})

	return messages, nil
}
