package controllers

import (
	"io"
	"io/ioutil"
	"net/http"
	"superMarket/wxapi"
	"time"
)

type WxController struct {
}

func validateSignature(r *http.Request) string {
	signature := r.FormValue("signature")
	timestamp := r.FormValue("timestamp")
	nonce := r.FormValue("nonce")
	echostr := r.FormValue("echostr")
	result := wxapi.ValidateSignature(signature, timestamp, nonce)
	if result {
		return echostr
	} else {
		return ""
	}
}

func (a *WxController) ListenMessage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	echostr := validateSignature(r)
	io.WriteString(w, echostr)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	msg, err := wxapi.ListenMessage(body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	respMsg := &wxapi.TextMessage{}
	respMsg.ToUserName.Text = msg.FromUserName.Text
	respMsg.FromUserName.Text = msg.ToUserName.Text
	respMsg.CreateTime = int(time.Now().Unix())
	respMsg.MsgType.Text = "text"
	respMsg.Content.Text = msg.Content.Text
	respMsg.MsgId = msg.MsgId
	result, err := respMsg.Marshal()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, result)
}
