package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"superMarket/repo"
	"superMarket/wxapi"
	"superMarket/wxapi/message"
	"time"
)

type JsapiTicket struct {
	AppId     string `json:"appId"`
	NonceStr  string `json:"nonceStr"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
}

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
		} else if eventType.Event.Text == "CLICK" {
			customMenuEvent := &message.CustomMenuEvent{}
			err := customMenuEvent.Unmarshal(body)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			result, err := customMenuEvent.Reply()
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			io.WriteString(w, result)
			return
		} else {

		}
	} else {
		io.WriteString(w, "没有匹配到相应的功能")
	}
}

func (a *WxController) QueryConfig(w http.ResponseWriter, r *http.Request) {
	nonceStr, err := createNonceStr()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	url := r.PostFormValue("url")
	jsapiTicket, err := queryJsapiTicket()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	signature := wxapi.CalculateSignature(nonceStr, jsapiTicket, timestamp, url)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	ticket := &JsapiTicket{AppId: wxapi.AppID, Timestamp: timestamp, NonceStr: nonceStr, Signature: signature}
	arrByte, err := json.Marshal(ticket)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(arrByte))
}

func createNonceStr() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	tempStr := fmt.Sprintf("%s", u)
	return tempStr, nil
}

func queryJsapiTicket() (string, error) {
	result, err := repo.GetKey("jsapi_ticket")
	if err != nil {
		return "", err
	}
	return result, nil
}
