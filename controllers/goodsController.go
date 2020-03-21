package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"superMarket/repo"
)

type GoodsController struct {
}

func (a *GoodsController) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	goodsCode := r.PostFormValue("goodsCode")
	goodsName := r.PostFormValue("goodsName")
	goodsBarCode := r.PostFormValue("goodsBarCode")
	categoryId, err := strconv.Atoi(r.PostFormValue("categoryId"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	goods := repo.NewGoods(0, goodsCode, goodsName, categoryId, goodsBarCode)
	err = goods.Create()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(goods)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}

func (a *GoodsController) Update(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	goodsCode := r.PostFormValue("goodsCode")
	goodsName := r.PostFormValue("goodsName")
	goodsBarCode := r.PostFormValue("goodsBarCode")
	goodsId, err := strconv.Atoi(r.PostFormValue("goodsId"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	categoryId, err := strconv.Atoi(r.PostFormValue("categoryId"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	goods := repo.NewGoods(goodsId, goodsCode, goodsName, categoryId, goodsBarCode)
	err = goods.Update()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(goods)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}

func (a *GoodsController) Delete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	goodsId, err := strconv.Atoi(r.PostFormValue("goodsId"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	goods := repo.NewGoods(goodsId, "", "", 0, "")
	err = goods.Delete()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(goods)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}

func (a *GoodsController) SelectById(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	goodsId, err := strconv.Atoi(r.PostFormValue("goodsId"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	goods := repo.NewGoods(goodsId, "", "", 0, "")
	err = goods.SelectById()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(goods)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}
