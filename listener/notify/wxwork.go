package notify

import (
	"encoding/json"
	"errors"
	"fmt"

	"supervisor-event-listener/event"
	"supervisor-event-listener/utils/httpclient"
)

type WorkWeixin struct{}

type Text struct {
	Content string `json:"content"`
	MentionedList []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

type WxMessage struct {
	MsgType string `json:"msgtype"`
	Text Text `json:"text"`
}

func (wx *WorkWeixin) Send(message event.Message) error {
	m := WxMessage{
		MsgType: "text",
		Text: Text{
			Content:message.String(),
			MentionedList: Conf.WorkWeixin.MentionedList,
		},
	}

	msg, err := json.Marshal(m)
	if err != nil {
		return err
	}

	timeout := 60
	response := httpclient.PostJson(Conf.WorkWeixin.Endpoint, string(msg), timeout)

	if response.StatusCode == 200 {
		return nil
	}
	errorMessage := fmt.Sprintf("workweixin执行失败#HTTP状态码-%d#HTTP-BODY-%s", response.StatusCode, response.Body)
	return errors.New(errorMessage)
}