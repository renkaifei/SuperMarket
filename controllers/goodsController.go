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
	goodsBarCode := r.PostFormValue("goodsBarCode")
	goodsName := r.PostFormValue("goodsName")
	goodsSpecification := r.PostFormValue("goodsSpecification")
	goodsDescription := r.PostFormValue("goodsDescription")
	goodsTradeMark := r.PostFormValue("goodsTradeMark")
	company := r.PostFormValue("company")
	goods := &repo.Goods{GoodsBarCode: goodsBarCode, GoodsName: goodsName, GoodsSpecification: goodsSpecification, GoodsDescription: goodsDescription, GoodsTradeMark: goodsTradeMark, Company: company}
	err := goods.Create()
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
	goodsBarCode := r.PostFormValue("goodsBarCode")
	goodsName := r.PostFormValue("goodsName")
	goodsSpecification := r.PostFormValue("goodsSpecification")
	goodsDescription := r.PostFormValue("goodsDescription")
	goodsTradeMark := r.PostFormValue("goodsTradeMark")
	company := r.PostFormValue("company")
	goodsId, err := strconv.Atoi(r.PostFormValue("goodsId"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	goods := &repo.Goods{GoodsId: goodsId, GoodsBarCode: goodsBarCode, GoodsName: goodsName, GoodsSpecification: goodsSpecification, GoodsDescription: goodsDescription, GoodsTradeMark: goodsTradeMark, Company: company}
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
	goodsId, err := strconv.Atoi(r.PostFormValue("goodsId"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	goods := &repo.Goods{GoodsId: goodsId}
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
	goodsId, err := strconv.Atoi(r.PostFormValue("goodsId"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	goods := &repo.Goods{GoodsId: goodsId}
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

func (a *GoodsController) SelectOnePage(w http.ResponseWriter, r *http.Request) {
	content := r.PostFormValue("content")
	pageIndex, err := strconv.Atoi(r.PostFormValue("pageIndex"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	pageSize, err := strconv.Atoi(r.PostFormValue("pageSize"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	goodses := &repo.Goodses{}
	goodses.Values = make([]*repo.Goods, 0)
	err = goodses.SelectOnePage(content, pageIndex, pageSize)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(goodses)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}

func (a *GoodsController) SelectByBarCode(w http.ResponseWriter, r *http.Request) {
	barCode := r.PostFormValue("barCode")
	goods := &repo.Goods{GoodsBarCode: barCode}
	err := goods.SelectByBarCode()
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
