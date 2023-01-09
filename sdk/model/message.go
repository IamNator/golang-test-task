package model

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

type (
	Message struct {
		Sender    string `json:"sender"`
		Receiver  string `json:"receiver"`
		Timestamp int64  `json:"timestamp"`
		Content   string `json:"content"`
	}
)

func (m Message) Byte() ([]byte, error) {
	return json.Marshal(m)
}

type (
	MessageRequest struct {
		Sender   string `json:"sender"`
		Receiver string `json:"receiver"`
		Content  string `json:"message"`
	}
)

func (m *MessageRequest) Bind(ctx *gin.Context) error {
	return ctx.BindJSON(m)
}

type MessageRequestValidateErr map[string]string

func (m MessageRequestValidateErr) Error() string {
	var s string
	for k, v := range m {
		s = k + " : " + v + "; "
	}
	return s
}
func (m MessageRequest) Validate() MessageRequestValidateErr {
	errs := make(MessageRequestValidateErr)
	if m.Sender == "" {
		errs["sender"] = "sender is required"
	}

	if m.Receiver == "" {
		errs["receiver"] = "receiver is required"
	}

	if m.Content == "" {
		errs["content"] = "content is required"
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

func (m MessageRequest) Data(sentAt time.Time) Message {
	return Message{
		Sender:    m.Sender,
		Receiver:  m.Receiver,
		Timestamp: sentAt.Unix(),
		Content:   m.Content,
	}
}
