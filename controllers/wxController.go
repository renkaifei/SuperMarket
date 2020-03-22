package controllers

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"superMarket/wxapi"
)

type WxController struct {
}

func (a *WxController) ValidateSignature(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	signature := r.FormValue("signature")
	timestamp := r.FormValue("timestamp")
	nonce := r.FormValue("nonce")
	echostr := r.FormValue("echostr")
	log.Println("signature:" + signature)
	log.Println("timestamp:" + timestamp)
	log.Println("nonce:" + nonce)
	log.Println("echostr:" + echostr)
	result := wxapi.ValidateSignature(signature, timestamp, nonce)
	if result {
		io.WriteString(w, echostr)
	} else {
		io.WriteString(w, "")
	}
}

func (a *WxController) FetchAccessToken() {

}

func (a *WxController) ListenMessage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	log.Println(string(body))
	msg, err := wxapi.ListenMessage(body)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, msg.URL.Text)
}
