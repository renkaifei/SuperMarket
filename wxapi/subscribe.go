package wxapi

import (
	"encoding/xml"
	"time"
)

type Subscribe struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   int
	MsgType      CDATA
	Event        CDATA
}

func (a *Subscribe) Marshal() (ret string, err error) {
	result, err := xml.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(result), err
}

func (a *Subscribe) Unmarshal(data []byte) error {
	err := xml.Unmarshal(data, a)
	return err
}

func (a *Subscribe) SubscribeReply() (reply string, err error) {
	message := &TextMessage{}
	message.FromUserName.Text = a.ToUserName.Text
	message.ToUserName.Text = a.FromUserName.Text
	message.CreateTime = int(time.Now().Unix())
	message.MsgType.Text = "text"
	message.Content.Text = "感谢您关注研飞超市，我们将竭尽全力为您提供最优质的本地化生活服务"
	result, err := xml.Marshal(message)
	return string(result), err
}
