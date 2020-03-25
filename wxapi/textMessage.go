package wxapi

import (
	"encoding/xml"
	"time"
)

type TextMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   int
	MsgType      CDATA
	Content      CDATA
	MsgId        int64
}

func (a *TextMessage) UnMarshal(data []byte) error {
	err := xml.Unmarshal(data, a)
	return err
}

func (a *TextMessage) ReplyTest() (ret string, err error) {
	v := &TextMessage{}
	v.ToUserName.Text = a.FromUserName.Text
	v.FromUserName.Text = a.ToUserName.Text
	v.CreateTime = int(time.Now().Unix())
	v.MsgType.Text = a.MsgType.Text
	v.Content.Text = a.Content.Text
	v.MsgId = a.MsgId
	result, err := xml.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
