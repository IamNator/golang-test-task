package worker

import (
	"github.com/gin-gonic/gin"
	"log"
	"twitch_chat_analysis/sdk/config"
	"twitch_chat_analysis/sdk/pkg/rabbitMQ"
	"twitch_chat_analysis/sdk/pkg/redis"
)

func Run() {

	if er := config.InitConfig(); er != nil {
		log.Fatalln(er.Error())
	}

	mqConn, er := rabbitMQ.NewConnection(config.GetConfig().RabbitMQURL)
	if er != nil {
		log.Fatalln(er.Error())
	}
	defer mqConn.Close()

	redisClient := redis.NewClient(config.GetConfig().RedisURL, "", 0)

	app := NewApp(mqConn, redisClient)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "ok")
	})

	if er := app.SubscribeToQueue(config.GetConfig().QueueName); er != nil {
		log.Fatalln(er)
	}

	r.Run(":" + config.GetConfig().WorkerPORT)
}
