package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"superMarket/repo"
)

type MerchanterController struct {
}

func (a *MerchanterController) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	openId := r.PostFormValue("openId")
	pwd := r.PostFormValue("pwd")
	merchanter := &repo.Merchanter{MerchanterOpenId: openId}
	err := merchanter.SelectByOpenId()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if merchanter.Pwd == pwd {
		goSessionId, err := repo.SetExpireKey(openId, openId, 60*20)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		merchanter.Pwd = ""
		result, err := json.Marshal(merchanter)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("goSessionId", goSessionId)
		io.WriteString(w, string(result))
	} else {
		http.Error(w, "用户密码错误", 500)
	}
}
