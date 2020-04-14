package message

import (
	"encoding/xml"
	"superMarket/repo"
	"superMarket/wxapi/menu"
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
	SessionId    string `xml:"-"`
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

func (a *TextMessage) Reply() (ret string, err error) {
	if a.Content.Text == "上货" {
		return a.UploadGoods()
	} else if a.Content.Text == "创建菜单" {
		return a.CreateMenu()
	} else if a.Content.Text == "删除菜单" {
		return a.DeleteMenu()
	} else {
		v := &TextMessage{}
		v.ToUserName.Text = a.FromUserName.Text
		v.FromUserName.Text = a.ToUserName.Text
		v.CreateTime = int(time.Now().Unix())
		v.MsgType.Text = a.MsgType.Text
		v.Content.Text = "没有匹配到对应的服务，请检查服务名称是否输入正确"
		data, err := xml.Marshal(v)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}
}

func (a *TextMessage) UploadGoods() (ret string, err error) {
	merchanter, err := getMerchanter(a.FromUserName.Text)
	if err != nil {
		return "", err
	}
	goSessionId, err := repo.SetExpireKey(a.FromUserName.Text, a.FromUserName.Text, 20*60)
	if err != nil {
		return "", err
	}
	v := &TextMessage{}
	v.ToUserName.Text = a.FromUserName.Text
	v.FromUserName.Text = a.ToUserName.Text
	v.CreateTime = int(time.Now().Unix())
	v.MsgType.Text = a.MsgType.Text
	if merchanter.IsAdmin == 1 {
		v.Content.Text = "<a href=\"http://www.daxuebaokao.cn/views/goods.html?goSessionId=" + goSessionId + "\">上传宝贝</a>"
	} else {
		v.Content.Text = "<a href=\"http://www.daxuebaokao.cn/views/merchantgoods.html?goSessionId=" + goSessionId + "&merchantId=" + merchanter.MerchantId + "\">上传宝贝</a>"
	}
	data, err := xml.Marshal(v)
	if err != nil {
		return err.Error(), nil
	}
	return string(data), nil
}

func (a *TextMessage) CreateMenu() (ret string, err error) {
	merchanter, err := getMerchanter(a.FromUserName.Text)
	if err != nil {
		return "", err
	}
	v := &TextMessage{}
	v.ToUserName.Text = a.FromUserName.Text
	v.FromUserName.Text = a.ToUserName.Text
	v.CreateTime = int(time.Now().Unix())
	v.MsgType.Text = a.MsgType.Text
	if merchanter.IsAdmin == 1 {
		result, err := menu.CreateMenu()
		if err != nil {
			return "", nil
		}
		v.Content.Text = result
	} else {
		v.Content.Text = "您没有更新菜单的权限"
	}
	data, err := xml.Marshal(v)
	return string(data), err
}

func (a *TextMessage) DeleteMenu() (ret string, err error) {
	merchanter, err := getMerchanter(a.FromUserName.Text)
	if err != nil {
		return "", err
	}
	v := &TextMessage{}
	v.ToUserName.Text = a.FromUserName.Text
	v.FromUserName.Text = a.ToUserName.Text
	v.CreateTime = int(time.Now().Unix())
	v.MsgType.Text = a.MsgType.Text
	if merchanter.IsAdmin == 1 {
		result, err := menu.DeleteMenu()
		if err != nil {
			return "", err
		}
		v.Content.Text = result
	} else {
		v.Content.Text = "您没有删除菜单的权限"
	}
	data, err := xml.Marshal(v)
	return string(data), err
}

func getMerchanter(openId string) (*repo.Merchanter, error) {
	merchanter := &repo.Merchanter{MerchanterOpenId: openId}
	err := merchanter.SelectByOpenId()
	return merchanter, err
}
