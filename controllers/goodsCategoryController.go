package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"superMarket/repo"
)

type GoodsCategoryController struct {
}

func (a *GoodsCategoryController) SelectByName(w http.ResponseWriter, r *http.Request) {
	categoryName := r.PostFormValue("categoryName")
	categorys := repo.GoodsCategorys{}
	items, err := categorys.SelectByName(categoryName)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}
