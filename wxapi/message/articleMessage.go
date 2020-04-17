package message

import (
	"encoding/xml"
)

type Article struct {
	Title       CDATA `xml:"item>Title"`
	Description CDATA `xml:"item>Description"`
	PicUrl      CDATA `xml:"item>PicUrl"`
	Url         CDATA `xml:"item>Url"`
}

type ArticleMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   int
	MsgType      CDATA
	ArticleCount int
	Articles     []*Article `xml:Articles>item,omitempty`
}
