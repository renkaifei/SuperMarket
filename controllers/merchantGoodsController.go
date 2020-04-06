package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"superMarket/repo"
)

type MerchantGoodsController struct {
}

func (a *MerchantGoodsController) SelectByMerchantIdAndGoodsId(w http.ResponseWriter, r *http.Request) {
	merchantId, err := strconv.Atoi(r.PostFormValue("merchantId"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	goodsId, err := strconv.Atoi(r.PostFormValue("goodsId"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	merchantGoods := repo.MerchantGoods{MerchantId: merchantId, GoodsId: goodsId}
	err = merchantGoods.SelectByMerchantIdAndGoodsId()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(merchantGoods)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}

func (a *MerchantGoodsController) Create(w http.ResponseWriter, r *http.Request) {
	merchantId, err := strconv.Atoi(r.PostFormValue("merchantId"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	goodsId, err := strconv.Atoi(r.PostFormValue("goodsId"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	price, err := strconv.ParseFloat(r.PostFormValue("price"), 64)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	discount, err := strconv.ParseFloat(r.PostFormValue("discount"), 64)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	merchantGoods := &repo.MerchantGoods{MerchantId: merchantId, GoodsId: goodsId, Price: price, Discount: discount}
	err = merchantGoods.Create()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(merchantGoods)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}

func (a *MerchantGoodsController) Update(w http.ResponseWriter, r *http.Request) {
	merchantGoodsId, err := strconv.Atoi(r.PostFormValue("merchantGoodsId"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	merchantId, err := strconv.Atoi(r.PostFormValue("merchantId"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	goodsId, err := strconv.Atoi(r.PostFormValue("goodsId"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	price, err := strconv.ParseFloat(r.PostFormValue("price"), 64)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	discount, err := strconv.ParseFloat(r.PostFormValue("discount"), 64)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	merchantGoods := &repo.MerchantGoods{MerchantGoodsId: merchantGoodsId, MerchantId: merchantId, GoodsId: goodsId, Price: price, Discount: discount}
	err = merchantGoods.Update()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(merchantGoods)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}

func (a *MerchantGoodsController) Delete(w http.ResponseWriter, r *http.Request) {
	merchantGoodsId, err := strconv.Atoi(r.PostFormValue("merchantGoodsId"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	merchantGoods := &repo.MerchantGoods{MerchantGoodsId: merchantGoodsId}
	err = merchantGoods.Delete()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(merchantGoods)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}

func (a *MerchantGoodsController) SelectPageByMerchantId(w http.ResponseWriter, r *http.Request) {
	content := r.PostFormValue("content")
	merchantId, err := strconv.Atoi(r.PostFormValue("merchantId"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
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
	merchantGoodses := &repo.MerchantGoodses{}
	err = merchantGoodses.SelectPageByMerchantId(content, merchantId, pageIndex, pageSize)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result, err := json.Marshal(merchantGoodses)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	io.WriteString(w, string(result))
}
