package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"superMarket/repo"
)

type MerchantController struct {
}

func (a *MerchantController) SelectById(w http.ResponseWriter, r *http.Request) {
	sessionId := r.Header.Get("goSessionId")
	value, err := repo.GetExpireKey(sessionId)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if value == "" {
		http.Error(w, "登陆过期，请重新登陆", 500)
		return
	}

	r.ParseForm()
	merchantId, err := strconv.Atoi(r.PostFormValue("merchantId"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	merchant := &repo.Merchant{MerchantId: merchantId}
	err = merchant.SelectById()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(merchant)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}
