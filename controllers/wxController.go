package controllers

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"superMarket/wxapi"
	"superMarket/wxapi/message"
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
	msgType := &message.MessageType{}
	err = msgType.UnMarshal(body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if msgType.MsgType.Text == "text" {
		textMsg := &message.TextMessage{}
		err := textMsg.UnMarshal(body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		str, err := textMsg.Reply()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		io.WriteString(w, str)
	} else if msgType.MsgType.Text == "event" {
		log.Println("msgType:event")
		eventType := &message.EventType{}
		err = eventType.Unmarshal(body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		if eventType.Event.Text == "subscribe" {
			subscribe := &message.Subscribe{}
			err := subscribe.Unmarshal(body)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			reply, err := subscribe.SubscribeReply()
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			io.WriteString(w, reply)
		} else {

		}
	} else {
		io.WriteString(w, "没有匹配到相应的功能")
	}
}
