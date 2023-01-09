package report

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"twitch_chat_analysis/sdk/config"
	"twitch_chat_analysis/sdk/pkg/redis"
)

func Run() {

	if er := config.InitConfig(); er != nil {
		log.Fatalln(er.Error())
	}

	redisClient := redis.NewClient(config.GetConfig().RedisURL, "", 0)
	app := NewApp(redisClient)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "ok")
	})

	r.GET("/message/list", func(c *gin.Context) {

		messages, er := app.GetMessages()
		if er != nil {
			c.JSONP(http.StatusUnprocessableEntity, er)
			return
		}

		c.JSON(200, messages)
	})

	r.Run(":" + config.GetConfig().ReportPORT)
}
