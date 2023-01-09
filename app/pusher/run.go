package pusher

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"twitch_chat_analysis/sdk/config"
	"twitch_chat_analysis/sdk/model"
	"twitch_chat_analysis/sdk/pkg/rabbitMQ"
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

	app := NewApp(mqConn)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "ok")
	})

	r.POST("/message", func(c *gin.Context) {

		var reqBody model.MessageRequest
		if er := reqBody.Bind(c); er != nil {
			c.JSON(http.StatusBadRequest, er)
			return
		}

		if er := reqBody.Validate(); er != nil {
			c.JSONP(http.StatusBadRequest, er)
			return
		}

		byteData, err := reqBody.Data(time.Now()).Byte()
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		if er := app.PushToQueue(config.GetConfig().QueueName, byteData); er != nil {
			c.JSON(http.StatusBadRequest, er)
			return
		}

		c.JSON(200, "pushed")
	})

	r.Run(":" + config.GetConfig().PusherPORT)
}
