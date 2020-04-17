package message

import (
	"encoding/xml"
	"log"
	"superMarket/repo"
	"time"
)

type CustomMenuEvent struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   int
	MsgType      CDATA
	Event        CDATA
	EventKey     CDATA
}

func (a *CustomMenuEvent) Unmarshal(data []byte) error {
	err := xml.Unmarshal(data, a)
	return err
}

func (a *CustomMenuEvent) Reply() (string, error) {
	if a.EventKey.Text == "todayRecommendation" {
		articles := &repo.Articles{PageIndex: 1, PageSize: 10}
		err := articles.SelectRecommend()
		if err != nil {
			log.Println("error:" + err.Error())
			return "", err
		}
		if len(articles.Values) > 0 {
			articleMessage := &ArticleMessage{}
			articleMessage.ToUserName.Text = a.FromUserName.Text
			articleMessage.FromUserName.Text = a.ToUserName.Text
			articleMessage.CreateTime = int(time.Now().Unix())
			articleMessage.MsgType.Text = "news"
			articleMessage.ArticleCount = len(articles.Values)
			for i := 0; i < articleMessage.ArticleCount; i++ {
				item := &Article{}
				item.Title.Text = articles.Values[i].Title
				item.Description.Text = articles.Values[i].Description
				item.PicUrl.Text = articles.Values[i].PicUrl
				item.Url.Text = articles.Values[i].Url
				articleMessage.Articles = append(articleMessage.Articles, item)
			}
			result, err := xml.Marshal(articleMessage)
			if err != nil {
				return "", err
			}
			return string(result), nil
		} else {
			textMessage := &TextMessage{}
			textMessage.ToUserName.Text = a.FromUserName.Text
			textMessage.FromUserName.Text = a.ToUserName.Text
			textMessage.CreateTime = int(time.Now().Unix())
			textMessage.MsgType.Text = "text"
			textMessage.Content.Text = "对不起，今天没有推荐的内容"
			result, err := xml.Marshal(textMessage)
			if err != nil {
				return "", err
			}
			return string(result), nil
		}
	} else {
		articleMessage := &ArticleMessage{}
		articleMessage.ToUserName.Text = a.FromUserName.Text
		articleMessage.FromUserName.Text = a.ToUserName.Text
		articleMessage.CreateTime = int(time.Now().Unix())
		articleMessage.MsgType.Text = "news"
		articleMessage.ArticleCount = 0
		data, err := xml.Marshal(articleMessage)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}
}
